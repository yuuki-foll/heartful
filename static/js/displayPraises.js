var cookies = document.cookie;
var cookiesArray = cookies.split(';');
console.log(cookiesArray)
for (var c of cookiesArray) {
    var cArray = c.split('=');
    console.log(cArray)
    if (cArray[0].trim(' ') === "comment") {
        var comment = cArray[1]
    }
    else if (cArray[0].trim(' ') === "name") {
        var user_name = cArray[1]
    }
}
console.log(comment);
var commentArray = comment.split(',')

// var commentList = []
// for (var i=0; i<commentArray.length; i++) {
//     commentList.push('<li>'+ user_name +commentArray[i] + '</li>');
// }

var heart_comment = document.getElementById("heart-comment")
for (let i=0; i < 10; i++) {
    if (i < commentArray.length) {
        heart_comment.insertAdjacentHTML('afterbegin', '<div class="comment"><li class="heart"><img src="/static/images/heart_pink.png" width="350px" height="350px"><p>'+user_name+commentArray[i]+'</p></li></div>')
    }
    else {
        heart_comment.insertAdjacentHTML('afterbegin', '<div class="comment"><li class="heart"><img src="/static/images/heart_pink.png" width="350px" height="350px"></li></div>')
    }
}
console.log(commentList)
document.getElementById("comment-list").innerHTML = commentList.join('');
