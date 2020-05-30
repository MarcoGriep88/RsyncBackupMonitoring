<template>

  <div id="app">
    <app-sidebar></app-sidebar>
    <!-- Content Wrapper -->
    <div  class="content-wrapper d-flex flex-column">

      <!-- Main Content -->
      <div class="content">
        <app-header></app-header>
        <app-about-overview></app-about-overview>
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

</style>
