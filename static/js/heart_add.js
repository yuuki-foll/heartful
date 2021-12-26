var list = document.getElementById("randomHeart");

for (let i = 0; i < 50; i++) {
    list.insertAdjacentHTML('afterbegin', '<li class="heart"><img src="/static/images/heart_pink.png" width="100px" height="100px"></li>');
    list.insertAdjacentHTML('afterbegin', '<li class="heart"><img src="/static/images/heart_red.png" width="100px" height="100px"></li>');
}