<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class FileInfo extends Model
{
    /**
     * The table associated with the model.
     *
     * @var string
     */
    protected $table = 'fileinfos';

    protected $fillable = array(
        'id',
        'fk_backupId',
        'file'
    );
}
