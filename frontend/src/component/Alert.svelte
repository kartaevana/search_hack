<script>
    export let message = "";
    export let onClose = () => {};
  </script>
  
  <style>
    .modal {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, 0.5);
      display: flex;
      justify-content: center;
      align-items: center;
    }
  
    .alert-box {
      background: white;
      padding: 20px;
      border-radius: 8px;
      box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
      text-align: center;
    }
  
    .close-button {
      margin-top: 20px;
    }
  </style>
  
  <div class="modal" on:click={onClose}>
    <div class="alert-box" on:click|stopPropagation>
      <p>{message}</p>
      <button class="close-button" on:click={onClose}>Закрыть</button>
    </div>
  </div>