let display = true;
window.addEventListener('click', () => {
    const btn_submit = document.getElementById("btn-submit");

    const formdata = new FormData()
    const radio = document.querySelectorAll(".radiobutton");
    radio.forEach(function(target) {
        target.addEventListener('click', () => {
            if (display) {
                document.getElementById("space-submit").innerHTML='<input id="btn-submit" type="submit" value="褒められにいく" class="submit">'
                display = false
            }
        })
    })
    btn_submit.addEventListener('click', (e) => {
        var audio = new Audio("/static/se/decide.mp3");
        audio.volume = 0.2;
        audio.play();

        e.preventDefault();

        const career = document.querySelector('input[name=career]:checked');
        formdata.append('career', career.value);
        console.log(career.value)
        console.log(formdata);
        // formdata.append('career', career.value);
        const url = "http://127.0.0.1:8080/get_praises"
        fetch(url, {
            method: 'POST',
            body: formdata,
        })
        .then(response => response.json())
        .then(data => {
            var comment = data.comment;
            document.cookie = "comment="+comment;
            console.log(comment);
            document.location.href = "http://127.0.0.1:8080/happy";
        })
        .catch((error) => {
            console.error(error);
        });
    });
});