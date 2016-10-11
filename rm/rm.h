#ifndef GO_RM_RM_H
#define GO_RM_RM_H

#include <stdlib.h>
#include "./redismodule.h"
#include "./callbacks.h"
#include "./wrapper.h"


int CreateCommandCallID(RedisModuleCtx *ctx,int id, const char *name,  const char *strflags, int firstkey, int lastkey, int keystep) {
  return RedisModule_CreateCommand(ctx, name, cb_cmd_func[id], strflags, firstkey, lastkey, keystep);
}

uintptr_t CreateDataTypeCallID(RedisModuleCtx* ctx,int id,const char* name,int encver){
    RedisModuleTypeLoadFunc rdb_load =cb_mt_rdb_load[id];
    RedisModuleTypeSaveFunc rdb_save =cb_mt_rdb_save[id];
    RedisModuleTypeRewriteFunc aof_rewrite =cb_mt_aof_rewrite[id];
    RedisModuleTypeDigestFunc digest =cb_mt_digest[id];
    RedisModuleTypeFreeFunc free =cb_mt_free[id];
    return (uintptr_t)CreateDataType(ctx,name,encver,rdb_load,rdb_save,aof_rewrite,digest,free);
}

#define LOG_DEBUG   "debug"
#define LOG_VERBOSE "verbose"
#define LOG_NOTICE  "notice"
#define LOG_WARNING "warning"

void CtxLog(RedisModuleCtx *ctx, int level, const char *fmt) {
  char *l;
  switch (level) {
    default:
    case 0:
      l = LOG_DEBUG;
      break;
    case 1:
      l = LOG_VERBOSE;
      break;
    case 2:
      l = LOG_NOTICE;
      break;
    case 3:
      l = LOG_WARNING;
      break;
  }
  RedisModule_Log(ctx, l, fmt);
}


int HashSetVar(RedisModuleKey *key, int flags,int argc, intptr_t argv[]){
    switch(argc){
case 0: return RedisModule_HashSet(key, flags);
case 1: return RedisModule_HashSet(key, flags,argv[0]);
case 2: return RedisModule_HashSet(key, flags,argv[0],argv[1]);
case 3: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2]);
case 4: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3]);
case 5: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4]);
case 6: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5]);
case 7: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6]);
case 8: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7]);
case 9: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8]);
case 10: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9]);
case 11: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10]);
case 12: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11]);
case 13: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12]);
case 14: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13]);
case 15: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14]);
case 16: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15]);
case 17: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15],argv[16]);
case 18: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15],argv[16],argv[17]);
case 19: return RedisModule_HashSet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15],argv[16],argv[17],argv[18]);

    default:
        return REDISMODULE_ERR;
    }
}

int HashGetVar(RedisModuleKey *key, int flags,int argc, intptr_t argv[]){
    switch(argc){
case 0: return RedisModule_HashGet(key, flags);
case 1: return RedisModule_HashGet(key, flags,argv[0]);
case 2: return RedisModule_HashGet(key, flags,argv[0],argv[1]);
case 3: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2]);
case 4: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3]);
case 5: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4]);
case 6: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5]);
case 7: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6]);
case 8: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7]);
case 9: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8]);
case 10: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9]);
case 11: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10]);
case 12: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11]);
case 13: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12]);
case 14: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13]);
case 15: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14]);
case 16: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15]);
case 17: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15],argv[16]);
case 18: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15],argv[16],argv[17]);
case 19: return RedisModule_HashGet(key, flags,argv[0],argv[1],argv[2],argv[3],argv[4],argv[5],argv[6],argv[7],argv[8],argv[9],argv[10],argv[11],argv[12],argv[13],argv[14],argv[15],argv[16],argv[17],argv[18]);

    default:
        return REDISMODULE_ERR;
    }
}

#endif
