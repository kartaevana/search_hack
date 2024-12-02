<script lang="ts">
    import "../../../app.css";
    function go_sign_up(){
        window.location.href = '/project/sign_up';
    }
    function go_sign_in(){
        window.location.href = '/project/sign_in';
    }
</script>

<header>
    <img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
</header>

<main>
    <div class="logo">
        <h1>Dream team</h1>
        <h2>Присоединяйтесь к нам!</h2>
        <div>
            <!-- <button on:click={save_note}>Сохранить</button> -->
            <input type="button" on:click={ go_sign_up } value="Регистрация" id="sign_up">
            <input type="button" on:click={ go_sign_in } value="Вход" id="sign_in">
        </div>
    </div>
</main>

<style lang="scss">
    main {
        display: flex;
        justify-content: center;
        /* align-content: center; */
            /* align-items: center; */
        .logo {
            height: 452px;
            /* background-color: #2c2c2c; */
            display: flex;
            justify-content: center;
            align-items: center;
            flex-direction: column;

            h1 {
                font-size: 86px;
            }

            h2 {
                font-size: 32px;
                color: rgba(255, 255, 255, 0.7);
            }

            input {
                margin-top: 20px;
                height: 32px;
                padding: 3px;
                border-radius: 8px;

            }
            #sign_up {
                    background-color: rgba(44, 44, 44, 1);
                    color: #ffffff;
                    width: 118px;
                }
            #sign_in{
                width: 55px;
            }
        }
    }
</style>