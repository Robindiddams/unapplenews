package unapplenews

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUrl(t *testing.T) {
	assert := assert.New(t)
	f, err := os.Open("testdata/article.html")
	assert.NoError(err)
	defer f.Close()
	url, err := findArticleURL(f)
	assert.EqualValues("https://arstechnica.com/tech-policy/2020/10/googles-supreme-court-faceoff-with-oracle-was-a-disaster-for-google/", url)
}
