<?php

namespace App\Http\Controllers;

use Laravel\Lumen\Routing\Controller as BaseController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

use App\BackupInfo;
use Response;

class Controller extends BaseController
{
    function __construct() {   }

    public function index() {
        $result = BackupInfo::all();
        return response()->json($result);
    }

    public function show_by_id($id) {
        $result = BackupInfo::where("id", $id)->first();
        return response()->json($result);
    }

    public function create(Request $request) {
        $checkElems = BackupInfo::where('Hostname', $request->Hostname)
                ->where('backupDate', $request->backupDate)
                ->where('backupType', $request->backupType)
                ->first();

        if ($checkElems!=null)
        {
            return "Already exists";
        }

        $obj = new BackupInfo;

        $obj->Hostname = $request->Hostname;
        $obj->backupDate = $request->backupDate;
        $obj->numerOfFiles = $request->numerOfFiles;

        $obj->created = $request->created;
        $obj->backupType = $request->backupType;

        $obj->save();

        return $obj;
    }

    public function clear(Request $request) {
        try {
            BackupInfo::where('Hostname', $request->Hostname)->delete();
            return "ok";
        }
        catch (Exception $e) {
            return "failed";
        }
    }
}
