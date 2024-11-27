package routes

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/handler"
	"github.com/gin-gonic/gin"
)

// InitRoutes 함수는 모든 API 라우트를 설정하고 반환합니다.
func InitRoutes() *gin.Engine {
	// Gin 엔진 생성
	router := gin.Default()

	// /posts 라우트 그룹
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("/", handler.CreatePost)    // 게시글 생성
		postRoutes.GET("/:id", handler.GetPost)     // 게시글 조회
		postRoutes.POST("/:id", handler.DeletePost) // 게시글 삭제
	}
	// /ai 라우트 그룹
	aiRoutes := router.Group("/ai")
	{
		aiRoutes.POST("/callback", handler.UpdatePostWithAI) // AI에서 보낸 데이터를 처리
	}

	return router
}
