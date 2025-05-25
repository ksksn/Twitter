document.querySelector('.twitter-container').addEventListener('click', function() {
    const icon = document.querySelector('i');
    const text = document.querySelector('h1');
    const body = document.body;
    
    const rect = icon.getBoundingClientRect();
    const centerX = rect.left + rect.width / 2;
    const centerY = rect.top + rect.height / 2;
    
    const circle = document.createElement('div');
    circle.className = 'expand';
    circle.style.top = centerY + 'px';
    circle.style.left = centerX + 'px';
    document.body.appendChild(circle);
    
    requestAnimationFrame(() => {
        circle.classList.add('active');
    });
    
    setTimeout(() => {
        text.style.opacity = '0';
    }, 500);
    
    setTimeout(() => {
        document.body.style.backgroundColor = 'white';
        icon.remove();
        text.remove();
        circle.remove();
        
        const loginContainer = document.querySelector('.login-container');
        loginContainer.classList.add('visible');
    }, 1500);
});

document.querySelector('.register-link').addEventListener('click', function(e) {
    e.preventDefault();
    
    const loginContainer = document.querySelector('.login-container');
    const registerContainer = document.querySelector('.register-container');
    
    loginContainer.classList.add('hidden');
    
    setTimeout(() => {
        registerContainer.classList.add('visible');
    }, 300);
});

// Обработчик для кнопок "Далее"
document.querySelectorAll('.next-button').forEach(button => {
    button.addEventListener('click', function() {
        const whiteScreen = document.createElement('div');
        whiteScreen.className = 'white-screen';
        document.body.appendChild(whiteScreen);
        
        // Скрываем все контейнеры
        document.querySelector('.login-container').style.display = 'none';
        document.querySelector('.register-container').style.display = 'none';
        
        // Активируем белый экран
        setTimeout(() => {
            whiteScreen.classList.add('active');
        }, 10);
    });
});

// Добавить новый обработчик для кнопки "Войти"
document.querySelector('.text-button').addEventListener('click', function() {
    const loginContainer = document.querySelector('.login-container');
    const registerContainer = document.querySelector('.register-container');
    
    registerContainer.classList.remove('visible');
    loginContainer.classList.remove('hidden');
    loginContainer.style.display = 'flex';
    
    setTimeout(() => {
        loginContainer.classList.add('visible');
    }, 100);
});