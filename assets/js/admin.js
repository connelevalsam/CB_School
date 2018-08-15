//================ Lecturer ======================
let dDiv     = document.querySelector('.default-div');
let cLect    = document.querySelector('.create');
let eLect    = document.querySelector('.edit');
let dLect    = document.querySelector('.delete');
let cLectDiv = document.querySelector('.add');
let eLectDiv = document.querySelector('.update');
let dLectDiv = document.querySelector('.remove');

function displayCdivs() {
    // body...

        dDiv.style.display     = 'none';
        eLectDiv.style.display = 'none';
        dLectDiv.style.display = 'none';
        var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz";
        var string_length = 20;
        var randomstring = '';
        for (var i=0; i<string_length; i++) {
            var rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum,rnum+1);

        document.creating.pword.value = randomstring;
        cLectDiv.style.display = 'block';
    }
}
function displayEdivs() {
    // body...
    if (eLectDiv.style.display === 'none') {
        var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz";
        var string_length = 20;
        var randomstring = '';
        for (var i=0; i<string_length; i++) {
            var rnum = Math.floor(Math.random() * chars.length);
            randomstring += chars.substring(rnum,rnum+1);
        }
        document.editing.pword.value = randomstring;
        dDiv.style.display     = 'none';
        cLectDiv.style.display = 'none';
        dLectDiv.style.display = 'none';
        eLectDiv.style.display = 'block';
    }
}
function displayDdivs() {
    // body...
    if (dLectDiv.style.display === 'none') {
        dDiv.style.display = 'none';
        cLectDiv.style.display = 'none';
        eLectDiv.style.display = 'none';
        dLectDiv.style.display = 'block';
    }
}

cLect.addEventListener('click', displayCdivs)
eLect.addEventListener('click', displayEdivs)
dLect.addEventListener('click', displayDdivs)

//
