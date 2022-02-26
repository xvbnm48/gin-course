# gin-course
golang gin course


Lets go to learning Gin !`


# how to use ? 

[//]: # (1. access to https://gin-videoss.herokuapp.com/view/videos , password auth is sakura_endo and username is sakura_endo)
1. login with username and password in route https://gin-videoss.herokuapp.com/login , username sakura_endo and password: sakura_endo with postman for get token
2. access to https://gin-videoss.herokuapp.com/view/videos 
3. if videos is not show, it means the video hasn't been uploaded yet
4. can upload with use postman and use this route http://gin-videos.herokuapp.com/api/posts , but use Bearer token for authorization in postman > authorization > bearer token , token that is got from the previous login
5. after filling authentication with Bearer (token) then fill json body like this 
 ```json
  {
    "title" : "how to become master golang",
    "description":"how to fast become master of golang under 1 years!",
    "url":"https://www.youtube.com/embed/JgW-i2QjgHQ",
    "author":{
      "first_name" :"sakura",
      "last_name":"endo",
      "age":19,
      "email":"sakuraendo@nogi.com"
    }
}
 ```
5. link video video links are obtained from embed YouTube
![](/templates/images/1.png?raw=true "Optional Title")
6. and click send at postman
7. if success create new videos, check on https://gin-videoss.herokuapp.com/view/videos , then the uploaded video will appear