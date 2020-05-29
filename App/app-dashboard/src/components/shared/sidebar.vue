<template>
<ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">

      <!-- Sidebar - Brand -->
      <a class="sidebar-brand d-flex align-items-center justify-content-center" href="/">
        <div class="sidebar-brand-icon rotate-n-15">
          <i class="fas fa-laugh-wink"></i>
        </div>
        <div class="sidebar-brand-text mx-3">Cockpit</div>
      </a>

      <!-- Divider -->
      <hr class="sidebar-divider my-0">

      <!-- Nav Item - Dashboard -->
      <li class="nav-item active">
        <a class="nav-link" href="/">
          <i class="fas fa-fw fa-tachometer-alt"></i>
          <span>Patch Management</span></a>
      </li>

      <li class="nav-item active">
        <a class="nav-link" href="/license">
          <i class="far fa-id-badge"></i>
          <span>License Dashboard</span></a>
      </li>

      <li class="nav-item active">
        <a class="nav-link" href="/backup">
          <i class="far fa-hdd"></i>
          <span>Managed Backup</span></a>
      </li>

      <li class="nav-item active">
        <a class="nav-link" href="/about">
          <i class="fas fa-info-circle"></i>
          <span>Info</span></a>
      </li>

      <!-- Divider -->
      <hr class="sidebar-divider d-none d-md-block">

      <!-- Sidebar Toggler (Sidebar) -->
      <div class="text-center d-none d-md-inline">
        <button class="rounded-circle border-0" id="sidebarToggle"></button>
      </div>

    </ul>
    <!-- End of Sidebar -->
</template>