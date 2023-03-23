package api

import (
	"chatgpt-api/config"
	"chatgpt-api/model"
	"chatgpt-api/restapi"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	defaultDefine = "You are a helpful assistant created by SJH not by OpenAI"
)

type ChatAPI struct{}

func (ca ChatAPI) RegisterRoute(r gin.IRouter) {
	r.POST("/chat", ca.Chat)
	r.GET("/chatHistory", ca.ChatHistory)
}

func (ChatAPI) setChatContext(sess sessions.Session, ctx []*model.GPTMessage) {
	sess.Set("chat-context", ctx)
}

func (ChatAPI) getChatContext(sess sessions.Session) []*model.GPTMessage {
	ctx := sess.Get("chat-context")
	if msgs, ok := ctx.([]*model.GPTMessage); ok {
		return msgs
	}
	return []*model.GPTMessage{{Role: model.RoleSystem, Content: defaultDefine}}
}

func (ca ChatAPI) Chat(c *gin.Context) {
	body := struct {
		Msg string `json:"msg" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	sess := sessions.Default(c)
	ctx := ca.getChatContext(sess)
	ctx = append(ctx, &model.GPTMessage{
		Role:    model.RoleUser,
		Content: body.Msg,
	})

	resp, err := restapi.ChatCompletion(&model.GPTRequest{
		Model:            model.ModelChat3_5,
		Messages:         ctx,
		Temperature:      0.5,
		FrequencyPenalty: 1,
		PresencePenalty:  1,
	}, config.APIKey())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if len(resp.Choices) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something wrong occurs"})
		return
	}
	msg := resp.Choices[0].Message

	ctx = append(ctx, msg)
	ca.setChatContext(sess, ctx)
	if err = sess.Save(); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"content": msg.Content,
	})
	return
}

func (ca ChatAPI) ChatHistory(c *gin.Context) {
	sess := sessions.Default(c)
	c.JSON(http.StatusOK, gin.H{
		"messages": ca.getChatContext(sess)[1:],
	})
}
