# gin-course
golang gin course


Lets go to learning Gin !


# how to use ? 

1. access to https://gin-videoss.herokuapp.com/view/videos , password auth is sakura_endo and username is sakura_endo
2. if videos is not show, it means the video hasn't been uploaded yet
3. can upload with use postman and use this route http://gin-videos.herokuapp.com/api/posts
4. fill out a form like this 
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
5. and click send at postman
6. if success create new videos, check on https://gin-videoss.herokuapp.com/view/videos , then the uploaded video will appear