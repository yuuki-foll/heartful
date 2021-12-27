let display = true;
const radio = document.querySelectorAll(".radiobutton");
radio.forEach(function(target) {
    target.addEventListener('click', () => {
        if (display) {
            document.getElementById("space-submit").innerHTML = '<input type="text" class="input_form" size="70" maxLength="50"  name="praise_comment" placeholder="コメントを入力してください（最大50文字） 例:さんって素敵ですね!" /><p><input type="submit" class="submit_button" value="コメントを送る!"></p>';
            display = false;
        }
    })
})