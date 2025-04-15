@echo off

REM Step 1: Create a new post
curl -X POST http://localhost:8080/posts -H "Content-Type: application/json" -d "{\"body\":\"This is a test post\"}" > response.json

echo Post created. Response saved to response.json

REM Step 2: Extract the ID of the created post
for /f "tokens=*" %%i in ('type response.json') do set response=%%i
for /f "tokens=2 delims=:," %%i in ('echo %response%') do set postId=%%i
set postId=%postId:~1,-1%

echo Extracted Post ID: %postId%

REM Step 3: Retrieve the created post
curl -X GET http://localhost:8080/posts/%postId% -H "Content-Type: application/json"

echo Retrieved post with ID %postId%

REM Step 4 (Optional): Delete the post
curl -X DELETE http://localhost:8080/posts/%postId% -H "Content-Type: application/json"

echo Deleted post with ID %postId%