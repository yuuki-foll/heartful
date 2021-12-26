window.addEventListener('DOMContentLoaded', () => {
    const btn_submit = document.getElementById("btn-submit");

    const formdata = new FormData()

    btn_submit.addEventListener('click', (e)=> {
        e.preventDefault();

        const career = document.querySelector('input[name=career]:checked');
        formdata.append('career', career.value);
        
        
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