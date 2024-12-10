<script lang="ts">
    import "../../app.css";
</script>
<header>
    <img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
</header>
<h1>Создание анкеты</h1>
<main>
    <form action="">
        <div class="question">
            <div><label for="name">ФИО</label></div>
            <div><input placeholder="Введите ваше ФИО" id="name"></div>
        </div>


        <div class="question">
            <div><label for="tg">Аккаунт в телеграмм</label></div>
            <div><input type="text" placeholder="@username"></div>
        </div>

        <div class="question">
            <div><label for="email">Почта</label></div>
            <div><input placeholder="1234@mail.ru" id="email" type="email"></div>
        </div>
        <div class="question">
            <div><label for="group">Академическая группа</label></div>
            <div><input placeholder="БПМ23-3" id="group" type="text"></div>
        </div>
        <div class="question">
            <div><label for="role">Роль</label></div>
            <div>
                <select id="roles" name="roles">
                    <option value="frontend">фронтенд</option>
                    <option value="backend">бэкенд</option>
                    <option value="designer">дизайнер</option>
                    <option value="ml">ML</option>
                    <!-- <option value="other">Другое</option> -->
                </select>
            </div>
        </div>
        <div class="question">
            <div><label for="description">О себе</label></div>
            <div>
                <textarea id="description" cols="113" rows="10"
                    placeholder="Я студент 2 курса мисис, начинающий фронтендер ищу команду для хакатона..."></textarea>
            </div>
        </div>
        <div class="question">

            <div><label for="achievements">Список достижений</label></div>
            <div>
                <textarea id="achievements" placeholder="2ое место на конкурсе 'Русский медвежонок'"></textarea>
            </div>
        </div>
        <div class="submit"><input type="submit" value="Создать анкету" id="submit"></div>
    </form>

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
            align-items: center;
            background-color: #1e1e1e;
            width: 977px;
            height: 704px;
            // margin: 10px;
            margin-bottom: 30px;



            .question {
                margin: 10px;

                input,
                select,
                textarea {
                    background-color: #1e1e1e;
                    width: 929px;
                    border-radius: 8px;
                    color: aliceblue;
                    height: 40px;
                }

                #description,
                #achievements {
                    height: 80px;

                }

            }

            #submit {

                background-color: rgba(245, 245, 245, 1);
                height: 190%;
                width: 310px;
                border-radius: 8px;
            }
        }
    }
</style>