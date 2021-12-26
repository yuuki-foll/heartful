/* ログイン画面に移動 */
const member = document.getElementById('logo_button');
member.addEventListener('click', function () {
    var move_url = ""
    move_url += window.location.protocol
    move_url += "//"
    move_url += window.location.hostname
    move_url += ":"
    move_url += window.location.port
    move_url += "/"
    console.log(move_url)
    window.location.href = move_url; // 通常の遷移
});