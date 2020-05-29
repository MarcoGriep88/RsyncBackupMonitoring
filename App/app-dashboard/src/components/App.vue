<template>

  <div id="app">
    <app-sidebar></app-sidebar>
    <!-- Content Wrapper -->
    <div  class="content-wrapper d-flex flex-column">

      <!-- Main Content -->
      <div class="content">
        <app-header></app-header>
        <app-upgradable-overview></app-upgradable-overview>
      </div>
      <footer class="sticky-footer bg-white">
        <div class="container my-auto">
          <div class="copyright text-center my-auto">
            <span>Copyright &copy; Marco Griep - 2020</span>
          </div>
        </div>
      </footer>
      <a class="scroll-to-top rounded" href="#page-top">
    <i class="fas fa-angle-up"></i>
  </a>
    </div>
  </div>
</template>

<script>
export default {
  
}
</script>

<style>
.loader {
  height: 5vmin;
  left: 50%;
  margin: 0 0 0 0;
  position: relative;
  top: 50%;
  width: 5vmin;
}
.loader:before {
  animation: blueRoute 4s linear infinite;
  background-color: #09f;
  content: "";
  height: 5vmin;
  left: 0;
  opacity: 0.5;
  position: absolute;
  top: 0;
  width: 5vmin;
}
.loader:after {
  animation: pinkRoute 4s linear infinite;
  background-color: #f09;
  content: "";
  height: 5vmin;
  left: 5vmin;
  opacity: 0.5;
  position: absolute;
  top: 5vmin;
  width: 5vmin;
}

@keyframes blueRoute {
  0% {
    transform: translate(0, 0) rotate(0);
  }
  25% {
    transform: translate(5vmin, 0) rotate(90deg);
  }
  50% {
    transform: translate(5vmin, 5vmin) rotate(180deg);
  }
  75% {
    transform: translate(0, 5vmin) rotate(270deg);
  }
  100% {
    transform: translate(0, 0) rotate(360deg);
  }
}
@keyframes pinkRoute {
  0% {
    transform: translate(0, 0);
  }
  25% {
    transform: translate(-5vmin, 0);
  }
  50% {
    transform: translate(-5vmin, -5vmin);
  }
  75% {
    transform: translate(0, -5vmin);
  }
  100% {
    transform: translate(0, 0);
  }
}
</style>
