package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"moderation_service/constants"
	"moderation_service/domain/models"
	"moderation_service/infrastructure/persistence"
	"moderation_service/infrastructure/repository"
	"net/http"
	"strings"
	"sync"
)

var (
	languages  = map[string]string{"ES": "languages/es.json", "EN": "languages/pt.json", "PT": "languages/en.json"}
	readWord   = map[string][]string{}
	moderation = []string{"putos", "puto", "puta", "muran", "perros", "mueran"}
)

type ModerationServiceImpl struct {
	moderationRepository repository.ModerationRepository
}

func InitModerationServiceImpl() *ModerationServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ModerationServiceImpl{
		moderationRepository: dbHelper.ModerationRepository,
	}
}

func (r *ModerationServiceImpl) ModerationProccess(moderation models.Moderation) (interface{}, models.Response) {

	if len(moderation.Input) > constants.MaxLenght {
		return nil, models.Response{
			Status:      http.StatusBadRequest,
			Description: "The text exceeds the limit",
		}
	}

	moderation.LangProccesed = DetectLanguage(moderation.Input)
	moderation.Ouput = Chunks(moderation.Input)

	return moderation, models.Response{
		Status: http.StatusOK,
	}
}

func Chunks(c string) string {

	a := strings.FieldsFunc(c, Split)

	for i, v := range a {
		a[i] = SplitChuck(strings.Fields(v))
	}

	return strings.Join(a, " ")
}

func Split(r rune) bool {
	return r == ':' || r == '.' || r == ',' || r == '?'
}

func SplitChuck(l []string) string {
	for i, v := range l {
		l[i] = ValidateWord(v)
	}

	return strings.Join(l, " ")
}

func ValidateWord(w string) string {

	for _, v := range moderation {
		var (
			in = strings.ToLower(w)
			to = strings.ToLower(v)
		)

		if strings.Contains(in, to) {
			d := LevenshteinDistance(in, to)
			if d == 0 {
				return fmt.Sprintf("%s%s", w, constants.SignPoint)
			}

			if d < 3 {
				return fmt.Sprintf("%s%s", w, constants.SignDoubt)
			}
		}
	}
	return w
}

func LevenshteinDistance(a, b string) int {

	if len(a) == 0 {
		return 1
	}

	if len(b) == 0 {
		return 1
	}

	if a == b {
		return 0
	}

	s1 := []rune(a)
	s2 := []rune(b)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	lenS1 := len(s1)
	lenS2 := len(s2)

	var x []uint16
	if lenS1+1 > constants.MinLengthThreshold {
		x = make([]uint16, lenS1+1)
	} else {
		x = make([]uint16, constants.MinLengthThreshold)
		x = x[:lenS1+1]
	}

	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}

	_ = x[lenS1]
	for i := 1; i <= lenS2; i++ {
		prev := uint16(i)
		for j := 1; j <= lenS1; j++ {
			current := x[j-1]
			if s2[i-1] != s1[j-1] {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return int(x[lenS1])
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func DetectLanguage(t string) []models.LangProccesed {
	var w sync.WaitGroup

	c := make(chan models.LangProccesed)
	w.Add(len(languages))

	go func() {
		w.Wait()
		close(c)
	}()

	for k, l := range languages {
		languageProccesor(c, &w, k, l, strings.Fields(t))
	}

	langProccesed := []models.LangProccesed{}
	for v := range c {
		langProccesed = append(langProccesed, v)
	}
	return langProccesed
}

func languageProccesor(c chan models.LangProccesed, w *sync.WaitGroup, k, path string, text []string) {
	var (
		words     []string
		proccesed models.LangProccesed
		total     float64
	)

	//if readWord[k] == nil {
	byteValue, _ := ioutil.ReadFile(path)
	json.Unmarshal([]byte(byteValue), &words)
	total = float64(len(text))
	readWord[k] = words
	//}

	go func() {
		defer w.Done()

		for _, t := range text {
			for _, wo := range words {
				if t == wo {
					proccesed.Calc++
					proccesed.Language = k
				}
			}
		}

		proccesed.Calc = proccesed.Calc / total
		c <- proccesed
	}()

}
