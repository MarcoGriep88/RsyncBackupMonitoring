<template>
      <main>
        <div class="container-fluid">
          <div class="row">
            <div class="col-xl-4 col-md-4 mb-4">
              <div class="card border-left-primary shadow h-100 py-2">
                <div class="card-body">
                  <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                      <div class="text-xs font-weight-bold text-primary text-uppercase mb-1">3. Party Patches</div>
                      <div class="loader" v-if="chocoLoader"></div>
                      <div class="h5 mb-0 font-weight-bold text-gray-800">{{ infos.length }}</div>
                    </div>
                    <div class="col-auto">
                      <i class="fab fa-firefox fa-2x text-gray-300"></i>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-xl-4 col-md-4 mb-4">
              <div class="card border-left-success shadow h-100 py-2">
                <div class="card-body">
                  <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                      <div class="text-xs font-weight-bold text-success text-uppercase mb-1">Offene Windows Updates</div>
                      <div class="loader" v-if="wsusLoader"></div>
                      <div class="h5 mb-0 font-weight-bold text-gray-800">{{ wsuses.length }}</div>
                    </div>
                    <div class="col-auto">
                      <i class="fab fa-windows fa-2x text-gray-300"></i>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Pending Requests Card Example -->
            <div class="col-xl-4 col-md-4 mb-4">
              <div class="card border-left-warning shadow h-100 py-2">
                <div class="card-body">
                  <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                      <div class="text-xs font-weight-bold text-warning text-uppercase mb-1">Treiber Updates</div>
                      <div class="h5 mb-0 font-weight-bold text-gray-800">{{ dums.length }}</div>
                    </div>
                    <div class="col-auto">
                      <i class="fas fa-laptop-medical fa-2x text-gray-300"></i>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Content Row -->

          <div class="row">

            <!-- Area Chart -->
            <div class="col-xl-7 col-lg-7">
              <div class="card shadow mb-4">
                <!-- Card Header - Dropdown -->
                <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between">
                  <h6 class="m-0 font-weight-bold text-primary">Third Party Updates</h6>
                  <div class="dropdown no-arrow">
                    <a class="dropdown-toggle" href="#" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <i class="fas fa-ellipsis-v fa-sm fa-fw text-gray-400"></i>
                    </a>
                    <div class="dropdown-menu dropdown-menu-right shadow animated--fade-in" aria-labelledby="dropdownMenuLink">
                      <div class="dropdown-header">Aktionen:</div>
                      <a class="dropdown-item" @click="fetchChocoData">Neu laden</a>
                    </div>
                  </div>
                </div>
                <!-- Card Body -->
                <div class="card-body">
                  <div class="chart-area">
                    <div class="loader" v-if="chocoLoader"></div>
                      <table class="table table-striped table-bordered table-responsive" id="myTable" style="width: 100%;">
                      <thead>
                        <tr>
                          <th>Hostname</th>
                          <th>Software</th>
                          <th>Installierte Version</th>
                          <th>Verfügbare Version</th>
                          <th>Gewichtung</th>
                        </tr>
                      </thead>
                      <tbody>
                      <tr v-for="u in infos" :key="u.id">
                        <th>{{u.Hostname}}</th>
                        <th>{{u.Software}}</th>
                        <th>{{u.LocalVersion}}</th>
                        <th>{{u.UpgradeVersion}}</th>
                        <th v-bind:class="{ 'text-warning': u.Weight > 10, 'text-danger': u.Weight > 100 }">{{u.Weight}}</th>
                      </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>

            <div class="col-xl-5 col-lg-5">
              <div class="card shadow mb-4">

                <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between">
                  <h6 class="m-0 font-weight-bold text-primary">Treiber Updates</h6>
                  <div class="dropdown no-arrow">
                    <a class="dropdown-toggle" href="#" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <i class="fas fa-ellipsis-v fa-sm fa-fw text-gray-400"></i>
                    </a>
                    <div class="dropdown-menu dropdown-menu-right shadow animated--fade-in" aria-labelledby="dropdownMenuLink">
                      <div class="dropdown-header">Dropdown Header:</div>
                      <a class="dropdown-item" @click="fetchDUMData">Neu laden</a>
                    </div>
                  </div>
                </div>

                <div class="card-body">
                  <div class="chart-pie pt-4 pb-2">
                    <div class="loader" v-if="dumLoader"></div>
                      <table class="table table-striped table-bordered table-responsive" id="myTabledum" style="width: 100%;">
                      <thead>
                        <tr>
                          <th>Hostname</th>
                          <th>Treiber</th>
                          <th>Verfügbare Version</th>
                        </tr>
                      </thead>
                      <tbody>
                      <tr v-for="u in dums" :key="u.id">
                        <th>{{u.Hostname}}</th>
                        <th>{{u.Driver}}</th>
                        <th>{{u.UpgradeVersion}}</th>
                      </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
  
          <!-- Content Row -->

          <div class="row">

            <!-- Area Chart -->
            <div class="col-12">
              <div class="card shadow mb-4">
                <!-- Card Header - Dropdown -->
                <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between">
                  <h6 class="m-0 font-weight-bold text-primary">Windows Updates</h6>
                  <div class="dropdown no-arrow">
                    <a class="dropdown-toggle" href="#" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <i class="fas fa-ellipsis-v fa-sm fa-fw text-gray-400"></i>
                    </a>
                    <div class="dropdown-menu dropdown-menu-right shadow animated--fade-in" aria-labelledby="dropdownMenuLink">
                      <div class="dropdown-header">Aktionen:</div>
                      <a class="dropdown-item" @click="fetchWSUSData">Neu laden</a>
                    </div>
                  </div>
                </div>
                <!-- Card Body -->
                <div class="card-body">
                  <div class="chart-area">
                    <div class="loader" v-if="wsusLoader"></div>
                      <table class="table table-striped table-bordered table-responsive" id="myTablewsus" style="width: 100%;">
                <thead>
                  <tr>
                    <th>Hostname</th>
                    <th>Betriebssystem</th>
                    <th>Update</th>
                    <th>KB</th>
                    <th>Empfohlene Aktion</th>
                    <th>Patch Datum</th>
                  </tr>
                </thead>
                <tbody>
                <tr v-for="w in wsuses" :key="w.id">
                  <th>{{w.Computername}}</th>
                  <th>{{w.OS}}</th>
                  <th>{{w.UpdateTitle}}</th>
                  <th>{{w.LegacyName}}</th>
                  <th>{{w.UpdateApprovalAction}}</th>
                  <th>{{w.CreationDate}}</th>
                </tr>
                </tbody>
              </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
</template>

<script>
export default {
  data() {
    return {
      info: {
        id: '',
        Hostname: '',
        Software: '',
        LocalVersion: '',
        UpgradeVersion: '',
        Weight: ''
      },
      infos: [],
      wsus: {
        id: '',
        UpdateTitle: '',
        LegacyName: '',
        SecurityBulletins: '',
        Computername: '',
        OS: '',
        IpAddress: '',
        UpdateInstallationStatus: '',
        UpdateApprovalAction: '',
        CreationDate: ''
      },
      wsuses: [],
      dum: {
        id: '',
        Hostname: '',
        Driver: '',
        UpgradeVersion: ''
      },
      dums: [],
      winfixedcount: 0,
      winnotFixedCount: 0,
      winPercentfixed: 35,
      winPercentNotfixed: 0,
      chocoLoader: true,
      wsusLoader: true,
      dumLoader: true
    };
  },
  created() {
    this.reload();
  },
  destroyed() {
    
  },
  methods: {
    submit() {
      console.log(this.user);
    },
    reload() {
      this.fetchChocoData();
      this.fetchDUMData();
      this.fetchWSUSData();
    },
    fetchChocoData() {
      this.$http.get('http://localhost:8090/upgrades')
          .then(response => {
            return response.json();
          })
          .then(data => {
            const resultArray = [];
            for (let key in data) {
              resultArray.push(data[key]);
            }
            this.infos = resultArray;
            this.chocoLoader = false;
            sleep(500).then(() => {
                $('#myTable').DataTable( {
                  "order": [[ 4, "desc"]],
                  "pageLength": 25,
                  "language": {
                      "url": "//cdn.datatables.net/plug-ins/9dcbecd42ad/i18n/German.json"
                  }
                });
            });
          });
      },
      fetchWSUSData() {
      this.$http.get('http://localhost:8095/')
          .then(response => {
            return response.json();
          })
          .then(data => {
            const resultArray = [];
            for (let key in data) {
              resultArray.push(data[key]);
            }
            this.wsuses = resultArray;
            this.wsusLoader = false;
            sleep(500).then(() => {
                $('#myTablewsus').DataTable( {
                  "order": [[ 5, "asc"]],
                  "language": {
                      "url": "//cdn.datatables.net/plug-ins/9dcbecd42ad/i18n/German.json"
                  }
                });
            });
          });
      },
      fetchDUMData() {
      this.$http.get('http://localhost:8100/')
          .then(response => {
            return response.json();
          })
          .then(data => {
            const resultArray = [];
            for (let key in data) {
              resultArray.push(data[key]);
            }
            this.dums = resultArray;
            this.dumLoader = false;
            sleep(500).then(() => {
                $('#myTabledum').DataTable( {
                  "pageLength": 25,
                  "language": {
                      "url": "//cdn.datatables.net/plug-ins/9dcbecd42ad/i18n/German.json"
                  }
                });
            });
          });
      }
  }
}

function sleep (time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}
</script>

<style>
.dataTables_wrapper {
  width: 100%;
}
</style>