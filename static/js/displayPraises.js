var cookies = document.cookie;
var cookiesArray = cookies.split(';');

for (var c of cookiesArray) {
    var cArray = c.split('=');

    if (cArray[0] === "comment") {
        var comment = cArray[1]
    }
}
console.log(comment);
var commentArray = comment.split(',')

var commentList = []
for (var i=0; i<commentArray.length; i++) {
    commentList.push('<li>'+ commentArray[i] + '</li>');
}
console.log(commentList)
document.getElementById("comment-list").innerHTML = commentList.join('');
