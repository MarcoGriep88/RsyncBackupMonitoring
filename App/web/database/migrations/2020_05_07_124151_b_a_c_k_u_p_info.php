<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class BackupInfo extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('backupinfos', function (Blueprint $table) {
            $table->increments('id');
            $table->text('Hostname');
            $table->text('backupDate');
            $table->text('numerOfFiles');
            $table->text('created');
            $table->text('backupType');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('backupinfos');
    }
}
