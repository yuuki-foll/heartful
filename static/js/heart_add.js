var list = document.getElementById("randomHeart");

for (let i = 0; i < 80; i++) {
    list.insertAdjacentHTML('afterbegin', '<li class="heart"><img src="/static/images/heart_pink.png" width="50px" height="50px"></li>');
    list.insertAdjacentHTML('afterbegin', '<li class="heart"><img src="/static/images/heart_red.png" width="50px" height="50px"></li>');
}