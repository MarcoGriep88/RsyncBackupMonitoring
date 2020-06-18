<template>
      <main>
        <div class="container-fluid">
          <div class="row">
            <div class="col-xl-12 col-md-12 mb-12">
              <div class="card border-left-primary shadow h-100 py-2">
                <div class="card-body">
                  <div class="row no-gutters align-items-center">
                    <div class="col mr-2">
                      <div class="text-xs font-weight-bold text-primary text-uppercase mb-1">Backup Reports</div>
                      <div class="loader" v-if="Loader"></div>
                      <div class="h5 mb-0 font-weight-bold text-gray-800">{{ infos.length }}</div>
                    </div>
                    <div class="col-auto">
                      <i class="fab fa-firefox fa-2x text-gray-300"></i>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <br/>

          <div class="row">

            <!-- Area Chart -->
            <div class="col-xl-12 col-md-12 mb-12">
              <div class="card shadow mb-4">
                <div class="card-header py-3 d-flex flex-row align-items-center justify-content-between">
                  <h6 class="m-0 font-weight-bold text-primary">Backups</h6>
                  <div class="dropdown no-arrow">
                    <a class="dropdown-toggle" href="#" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <i class="fas fa-ellipsis-v fa-sm fa-fw text-gray-400"></i>
                    </a>
                    <div class="dropdown-menu dropdown-menu-right shadow animated--fade-in" aria-labelledby="dropdownMenuLink">
                      <div class="dropdown-header">Actions:</div>
                      <a class="dropdown-item" @click="fetchData">Reload</a>
                    </div>
                  </div>
                </div>
                <!-- Card Body -->
                <div class="card-body">
                  <div class="chart-area">
                    <div class="loader" v-if="Loader"></div>
                      <table class="table table-striped table-bordered table-responsive" id="myTable" style="width: 100%;">
                      <thead>
                        <tr>
                          <th>Hostname</th>
                          <th>Backup Date</th>
                          <th>Number of Backup Files</th>
                          <th>New Files created</th>
                          <th>Backup Description</th>
                        </tr>
                      </thead>
                      <tbody>
                      <tr v-for="u in infos" :key="u.id">
                        <td><router-link :to="{ path: '/details/' + u.id }"><a>{{ u.Hostname }}</a></router-link></td>
                        <td>{{ u.backupDate }}</td>
                        <td>{{ u.numerOfFiles }}</td>
                        <td>{{ u.created }}</td>
                        <td>{{ u.backupType }}</td>
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
        backupDate: '',
        numerOfFiles: '',
        created: '',
        backupType: ''
      },
      infos: [],
      Loader: true,
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
      this.fetchData();
    },
    fetchData() {
      this.$http.get('http://localhost:15000/')
          .then(response => {
            return response.json();
          })
          .then(data => {
            const resultArray = [];
            for (let key in data) {
              resultArray.push(data[key]);
            }
            this.infos = resultArray;
            this.Loader = false;
            sleep(500).then(() => {
                $('#myTable').DataTable( {
                  "order": [[ 1, "Backup Date"]],
                  "pageLength": 500,
                  "language": {
                      "url": "//cdn.datatables.net/plug-ins/9dcbecd42ad/i18n/German.json"
                  }
                });
            });
          });
      },
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
