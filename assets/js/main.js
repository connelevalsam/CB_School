//hamburger
    let hamburger = document.querySelector('#hamburger'),
    sidebar       = document.querySelector('#sidebar'),
    content       = document.querySelector('.contents'),
    sidebarWidth  = sidebar.getBoundingClientRect().width;

    //functions
    function toggleHamburger(event) {
    	event.preventDefault();
    	sidebar.classList.toggle('transform-off');

    hamburger.style.transform = hamburger.style.transform ? '' : 'translate3d(-' + sidebarWidth + 'px, 0, 0)';
    }

    function bodyClick() {
    	hamburger.style.transform = '';
    	sidebar.classList.add('transform-off');
    }
    //events listeners
    hamburger.addEventListener('click', toggleHamburger);
    content.addEventListener('click', bodyClick);