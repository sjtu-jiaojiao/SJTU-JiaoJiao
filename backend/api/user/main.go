package main

import (
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/user", getUser)
	return router
}

/**
 * @api {get} /user GetUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission guest/user/admin
 * @apiName GetUser
 * @apiDescription Get user info, return JSON list, only exist column will be return!
 *
 * @apiParam {int32} (None - Parameter) userId user id
 * @apiParam {int32} (Admin - Parameter) [userId] filter user id
 * @apiParam {string} (Admin - Parameter) [userName] filter user name
 * @apiParam {string} (Admin - Parameter) [telephone] filter user telephone
 * @apiParam {string} (Admin - Parameter) [studentId] filter student id
 * @apiParam {string} (Admin - Parameter) [studentName] filter student name
 * @apiParam {int32} (Admin - Parameter) limit=10 filter limit
 * @apiParam {int32} (Admin - Parameter) offset=0 filter offset
 *
 * @apiSuccess (Guest - Success 200) {int32} userId user id
 * @apiSuccess (Guest - Success 200) {string} userName user name
 * @apiSuccess (Guest - Success 200) {string} avatarId user avatar id
 * @apiSuccess (Guest - Success 200) {string} telephone user telephone
 *
 * @apiSuccess (User/Admin - Success 200) {int32} userId user id
 * @apiSuccess (User/Admin - Success 200) {string} userName user name
 * @apiSuccess (User/Admin - Success 200) {string} avatarId user avatar id
 * @apiSuccess (User/Admin - Success 200) {string} telephone user telephone
 * @apiSuccess (User/Admin - Success 200) {string} studentId student id
 * @apiSuccess (User/Admin - Success 200) {string} studentName student name
 *
 * @apiSuccess (Error 500) UserServiceDown User service down
 */
func getUser(c *gin.Context) {

}

func main() {
	utils.RunWebService("user", setupRouter())
}
