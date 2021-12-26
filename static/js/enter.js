var name_btn = document.getElementById("name-button");

name_btn.addEventListener('click', () => {
    var name = document.getElementById("input-name");
    const user_name =  name.value;
    const output = "ようこそ!"+ user_name +"さん";
    document.cookie = "name="+user_name;
    document.getElementById("output-name").innerText = output;
});