<script lang="ts">
    import "../../../app.css";
</script>

<header>
    <img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
</header>
<h1>Вход</h1>
<main>
    <form action="">
        <div class="question">
            <div><label for="email">Почта</label></div>
            <div><input placeholder="1234@mail.ru" id="email" type="email"></div>
        </div>
        <div class="question">
            <div><label for="password">Пароль</label></div>
            <div><input placeholder="Введите пароль" id="password" type="password"></div>
        </div>

        <div class="submit"><input type="submit" value="Войти" id="submit"></div>
    </form>
    <!-- <div><a href="">Забыли пароль?</a></div> -->

</main>

<style lang="scss">
    h1 {
        display: flex;
        justify-content: center;
        margin: 30px;
        font-size: 70px;
    }

    main {
        display: flex;
        justify-content: center;

        form {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            background-color: #1e1e1e;
            width: 578px;
            height: 322px;

            input {
                background-color: #1e1e1e;
                width: 530px;
                border-radius: 8px;
                color: aliceblue;
                height: 40px;
            }

            .question {
                margin: 10px;
            }

            #submit {

                background-color: rgba(245, 245, 245, 1);
                height: 190%;
                width: 310px;
                border-radius: 8px;
                color: black;
            }
        }
    }
</style>