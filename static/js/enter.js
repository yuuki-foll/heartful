var name_btn = document.getElementById("name-button");

name_btn.addEventListener('click', () => {
    var name = document.getElementById("input-name");
    const user_name =  name.value;
    const output = "ようこそ!"+ user_name +"さん";
    document.cookie = "name="+user_name;
    document.getElementById("output-name").innerText = output;
    document.getElementById("link1").innerHTML = '<a id="link1-btn" class="link-btn">褒めたい人はこちら</a>';
    document.getElementById("link2").innerHTML = '<a id="link2-btn" class="link-btn">褒められたい人はこちら</a>';
    add_url();
});

function add_url() {
    const link1 = document.getElementById("link1-btn");
    link1.href = "http://127.0.0.1:8080/praise";
    const link2 = document.getElementById("link2-btn");
    link2.href = "http://127.0.0.1:8080/user_info";
}