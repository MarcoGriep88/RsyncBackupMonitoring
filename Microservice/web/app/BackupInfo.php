<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class BackupInfo extends Model
{
    /**
     * The table associated with the model.
     *
     * @var string
     */
    protected $table = 'backupinfos';

    protected $fillable = array(
        'id',
        'Hostname',
        'backupDate',
        'numerOfFiles',
        'created',
        'backupType'
    );
}
