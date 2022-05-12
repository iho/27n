package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"27n/client"
	"27n/client/matches"

	"github.com/gin-gonic/gin"
)

var (
	storage                        MatchStorage
	timeoutErrorsPercantage        = 15
	internalServerErrorsPercantage = 15
	timeoutSeconds                 = 1
)

func init() {
	storage = MatchStorage{
		matches: make(map[int64]*matches.GetMatchesMatchIDOKBody),
	}
}

func main() {
	router := gin.Default()
	router.Use(ProblematicMiddleware())
	router.GET("/matches", getAllMatches)
	router.GET("/matches/:match_id", getAllMatchByID)
	router.GET("/", index)

	router.Run("localhost:8080")
}

func index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func getAllMatches(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.GetAllMatches())
}

func getAllMatchByID(c *gin.Context) {

	matchID, err := strconv.ParseInt(c.Param("match_id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad match id number"})
		return
	}

	match, err := storage.GetMatch(matchID)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "match not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, match)
}

type MatchStorage struct {
	matches     map[int64]*matches.GetMatchesMatchIDOKBody
	matchesLock sync.RWMutex
}

func (s *MatchStorage) GetMatch(ID int64) (*matches.GetMatchesMatchIDOKBody, error) {
	s.matchesLock.RLock()
	match, ok := s.matches[ID]
	s.matchesLock.RUnlock()
	if ok {
		return match, nil
	}
	params := matches.NewGetMatchesMatchIDParams().WithDefaults().WithMatchID(ID)
	response, err := client.Default.Matches.GetMatchesMatchID(params)
	if err != nil {
		return nil, err
	}
	match = response.GetPayload()

	s.matchesLock.Lock()
	s.matches[ID] = match
	s.matchesLock.Unlock()

	return match, nil
}

func (s *MatchStorage) GetAllMatches() []*matches.GetMatchesMatchIDOKBody {
	s.matchesLock.RLock()
	defer s.matchesLock.RUnlock()
	result := make([]*matches.GetMatchesMatchIDOKBody, 0, len(s.matches))

	for _, match := range s.matches {
		result = append(result, match)
	}
	return result
}

// ProblematicMiddleware makes handles fail sometimes
func ProblematicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch i := rand.Intn(100); {
		case i >= (100 - timeoutErrorsPercantage):
			c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{"message": "internal error"})
		case i >= (100 - timeoutErrorsPercantage - internalServerErrorsPercantage):
			time.Sleep(time.Duration(timeoutSeconds) * time.Second)
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"message": "time out error"})
		default:
			c.Next()
		}
	}
}
