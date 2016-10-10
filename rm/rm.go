package rm

//#include "./rm.h"
import "C"
import (
    "unsafe"
    "fmt"
)
/* ---------------- Defines common between core and modules --------------- */

/* Error status return values. */
const OK = C.REDISMODULE_OK
const ERR = C.REDISMODULE_ERR

/* API versions. */
const APIVER_1 = C.REDISMODULE_APIVER_1

/* API flags and constants */
const READ = C.REDISMODULE_READ
const WRITE = C.REDISMODULE_WRITE

const LIST_HEAD = C.REDISMODULE_LIST_HEAD
const LIST_TAIL = C.REDISMODULE_LIST_TAIL

/* Key types. */
const KEYTYPE_EMPTY = C.REDISMODULE_KEYTYPE_EMPTY
const KEYTYPE_STRING = C.REDISMODULE_KEYTYPE_STRING
const KEYTYPE_LIST = C.REDISMODULE_KEYTYPE_LIST
const KEYTYPE_HASH = C.REDISMODULE_KEYTYPE_HASH
const KEYTYPE_SET = C.REDISMODULE_KEYTYPE_SET
const KEYTYPE_ZSET = C.REDISMODULE_KEYTYPE_ZSET
const KEYTYPE_MODULE = C.REDISMODULE_KEYTYPE_MODULE

/* Reply types. */
const REPLY_UNKNOWN = C.REDISMODULE_REPLY_UNKNOWN
const REPLY_STRING = C.REDISMODULE_REPLY_STRING
const REPLY_ERROR = C.REDISMODULE_REPLY_ERROR
const REPLY_INTEGER = C.REDISMODULE_REPLY_INTEGER
const REPLY_ARRAY = C.REDISMODULE_REPLY_ARRAY
const REPLY_NULL = C.REDISMODULE_REPLY_NULL

/* Postponed array length. */
const POSTPONED_ARRAY_LEN = C.REDISMODULE_POSTPONED_ARRAY_LEN

/* Expire */
const NO_EXPIRE = C.REDISMODULE_NO_EXPIRE

/* Sorted set API flags. */
const ZADD_XX = C.REDISMODULE_ZADD_XX
const ZADD_NX = C.REDISMODULE_ZADD_NX
const ZADD_ADDED = C.REDISMODULE_ZADD_ADDED
const ZADD_UPDATED = C.REDISMODULE_ZADD_UPDATED
const ZADD_NOP = C.REDISMODULE_ZADD_NOP

/* Hash API flags. */
const HASH_NONE = C.REDISMODULE_HASH_NONE
const HASH_NX = C.REDISMODULE_HASH_NX
const HASH_XX = C.REDISMODULE_HASH_XX
const HASH_CFIELDS = C.REDISMODULE_HASH_CFIELDS
const HASH_EXISTS = C.REDISMODULE_HASH_EXISTS

/* A special pointer that we can use between the core and the module to signal
 * field deletion, and that is impossible to be a valid pointer. */
//const HASH_DELETE = C.REDISMODULE_HASH_DELETE

/* Error messages. */
const ERRORMSG_WRONGTYPE = C.REDISMODULE_ERRORMSG_WRONGTYPE

//const POSITIVE_INFINITE = C.REDISMODULE_POSITIVE_INFINITE
//const NEGATIVE_INFINITE = C.REDISMODULE_NEGATIVE_INFINITE

/* ------------------------- End of common defines ------------------------ */


// Use like malloc(). Memory allocated with this function is reported in
// Redis INFO memory, used for keys eviction according to maxmemory settings
// and in general is taken into account as memory allocated by Redis.
// You should avoid to use malloc().
// void *RM_Alloc(size_t bytes);
func Alloc(bytes int) (unsafe.Pointer) {
    return unsafe.Pointer(C.Alloc(C.size_t(bytes)))
}

// Use like realloc() for memory obtained with `RedisModule_Alloc()`.
// void* RM_Realloc(void *ptr, size_t bytes);
func Realloc(ptr unsafe.Pointer, bytes int) (unsafe.Pointer) {
    return unsafe.Pointer(C.Realloc(ptr, C.size_t(bytes)))
}

// Use like free() for memory obtained by `RedisModule_Alloc()` and
// `RedisModule_Realloc()`. However you should never try to free with
// `RedisModule_Free()` memory allocated with malloc() inside your module.
// void RM_Free(void *ptr);
func Free(ptr unsafe.Pointer) () {
    C.Free(ptr)
}

// Like strdup() but returns memory allocated with `RedisModule_Alloc()`.
// char *RM_Strdup(const char *str);
func Strdup(str unsafe.Pointer) (unsafe.Pointer) {
    return unsafe.Pointer(C.Strdup((*C.char)(str)))
}

// Lookup the requested module API and store the function pointer into the
// target pointer. The function returns `REDISMODULE_ERR` if there is no such
// named API, otherwise `REDISMODULE_OK`.
//
// This function is not meant to be used by modules developer, it is only
// used implicitly by including redismodule.h.
// int RM_GetApi(const char *funcname, void **targetPtrPtr);
//func GetApi(funcname string,targetPtrPtr /* TODO void** */unsafe.Pointer)(int){return int(C.GetApi(funcname,targetPtrPtr))}


// =============================================================================
// ========================== Context functions
// =============================================================================

// int `RedisModule_OnLoad(RedisModuleCtx` *ctx) {
//          // some code here ...
//          BalancedTreeType = `RM_CreateDataType(`...);
//      }
// moduleType *RM_CreateDataType(RedisModuleCtx *ctx, const char *name, int encver, moduleTypeLoadFunc rdb_load, moduleTypeSaveFunc rdb_save, moduleTypeRewriteFunc aof_rewrite, moduleTypeDigestFunc digest, moduleTypeFreeFunc free);
// NOTE
//func (ctx Ctx)CreateDataType(name string,encver int,rdb_load RedisModuleTypeLoadFunc,rdb_save RedisModuleTypeSaveFunc,aof_rewrite RedisModuleTypeRewriteFunc,digest RedisModuleTypeDigestFunc,free RedisModuleTypeFreeFunc)(/* TODO RedisModuleType* */unsafe.Pointer){return /* TODO RedisModuleType* */unsafe.Pointer(C.CreateDataType(ctx.ptr(),name,encver,rdb_load,rdb_save,aof_rewrite,digest,free))}


// Return heap allocated memory that will be freed automatically when the
// module callback function returns. Mostly suitable for small allocations
// that are short living and must be released when the callback returns
// anyway. The returned memory is aligned to the architecture word size
// if at least word size bytes are requested, otherwise it is just
// aligned to the next power of two, so for example a 3 bytes request is
// 4 bytes aligned while a 2 bytes request is 2 bytes aligned.
//
// There is no realloc style function since when this is needed to use the
// pool allocator is not a good idea.
//
// The function returns NULL if `bytes` is 0.
// void *RM_PoolAlloc(RedisModuleCtx *ctx, size_t bytes);
func (ctx Ctx)PoolAlloc(bytes int) (unsafe.Pointer) {
    return unsafe.Pointer(C.PoolAlloc(ctx.ptr(), C.size_t(bytes)))
}

// Return non-zero if a module command, that was declared with the
// flag "getkeys-api", is called in a special way to get the keys positions
// and not to get executed. Otherwise zero is returned.
// int RM_IsKeysPositionRequest(RedisModuleCtx *ctx);
func (ctx Ctx)IsKeysPositionRequest() (int) {
    return int(C.IsKeysPositionRequest(ctx.ptr()))
}

// When a module command is called in order to obtain the position of
// keys, since it was flagged as "getkeys-api" during the registration,
// the command implementation checks for this special call using the
// `RedisModule_IsKeysPositionRequest()` API and uses this function in
// order to report keys, like in the following example:
//
//  if (`RedisModule_IsKeysPositionRequest(ctx))` {
//      `RedisModule_KeyAtPos(ctx`,1);
//      `RedisModule_KeyAtPos(ctx`,2);
//  }
//
//  Note: in the example below the get keys API would not be needed since
//  keys are at fixed positions. This interface is only used for commands
//  with a more complex structure.
// void RM_KeyAtPos(RedisModuleCtx *ctx, int pos);
func (ctx Ctx)KeyAtPos(pos int) () {
    C.KeyAtPos(ctx.ptr(), C.int(pos))
}

// And is supposed to always return `REDISMODULE_OK`.
//
// The set of flags 'strflags' specify the behavior of the command, and should
// be passed as a C string compoesd of space separated words, like for
// example "write deny-oom". The set of flags are:
//
// * **"write"**:     The command may modify the data set (it may also read
//                    from it).
// * **"readonly"**:  The command returns data from keys but never writes.
// * **"admin"**:     The command is an administrative command (may change
//                    replication or perform similar tasks).
// * **"deny-oom"**:  The command may use additional memory and should be
//                    denied during out of memory conditions.
// * **"deny-script"**:   Don't allow this command in Lua scripts.
// * **"allow-loading"**: Allow this command while the server is loading data.
//                        Only commands not interacting with the data set
//                        should be allowed to run in this mode. If not sure
//                        don't use this flag.
// * **"pubsub"**:    The command publishes things on Pub/Sub channels.
// * **"random"**:    The command may have different outputs even starting
//                    from the same input arguments and key values.
// * **"allow-stale"**: The command is allowed to run on slaves that don't
//                      serve stale data. Don't use if you don't know what
//                      this means.
// * **"no-monitor"**: Don't propoagate the command on monitor. Use this if
//                     the command has sensible data among the arguments.
// * **"fast"**:      The command time complexity is not greater
//                    than O(log(N)) where N is the size of the collection or
//                    anything else representing the normal scalability
//                    issue with the command.
// * **"getkeys-api"**: The command implements the interface to return
//                      the arguments that are keys. Used when start/stop/step
//                      is not enough because of the command syntax.
// * **"no-cluster"**: The command should not register in Redis Cluster
//                     since is not designed to work with it because, for
//                     example, is unable to report the position of the
//                     keys, programmatically creates key names, or any
//                     other reason.
// int RM_CreateCommand(RedisModuleCtx *ctx, const char *name, RedisModuleCmdFunc cmdfunc, const char *strflags, int firstkey, int lastkey, int keystep);
//func (ctx Ctx)CreateCommand(name string, cmdfunc CmdFunc, strflags string, firstkey int, lastkey int, keystep int) (int) {
//    return int(C.CreateCommand(ctx.ptr(), name, cmdfunc, strflags, firstkey, lastkey, keystep))
//}

// Called by `RM_Init()` to setup the `ctx->module` structure.
//
// This is an internal function, Redis modules developers don't need
// to use it.
// void RM_SetModuleAttribs(RedisModuleCtx *ctx, const char *name, int ver, int apiver);
func (ctx Ctx)SetModuleAttribs(name string, ver int, apiver int) () {
    C.SetModuleAttribs(ctx.ptr(), C.CString(name), C.int(ver), C.int(apiver))
}

// Enable automatic memory management. See API.md for more information.
//
// The function must be called as the first function of a command implementation
// that wants to use automatic memory.
// void RM_AutoMemory(RedisModuleCtx *ctx);
func (ctx Ctx)AutoMemory() () {
    C.AutoMemory(ctx.ptr())
}

// Create a new module string object. The returned string must be freed
// with `RedisModule_FreeString()`, unless automatic memory is enabled.
//
// The string is created by copying the `len` bytes starting
// at `ptr`. No reference is retained to the passed buffer.
// RedisModuleString *RM_CreateString(RedisModuleCtx *ctx, const char *ptr, size_t len);
func (ctx Ctx)CreateString(ptr string, len int) (String) {
    c := C.CString(ptr)
    defer C.free(c);
    return CreateString(unsafe.Pointer(C.CreateString(ctx.ptr(), c, C.size_t(len))))
}

// Like `RedisModule_CreatString()`, but creates a string starting from a long long
// integer instead of taking a buffer and its length.
//
// The returned string must be released with `RedisModule_FreeString()` or by
// enabling automatic memory management.
// RedisModuleString *RM_CreateStringFromLongLong(RedisModuleCtx *ctx, long long ll);
func (ctx Ctx)CreateStringFromLongLong(ll int64) (String) {
    return CreateString(unsafe.Pointer(C.CreateStringFromLongLong(ctx.ptr(), C.longlong(ll))))
}

// Like `RedisModule_CreatString()`, but creates a string starting from an existing
// RedisModuleString.
//
// The returned string must be released with `RedisModule_FreeString()` or by
// enabling automatic memory management.
// RedisModuleString *RM_CreateStringFromString(RedisModuleCtx *ctx, const RedisModuleString *str);
func (ctx Ctx)CreateStringFromString(str String) (String) {
    return CreateString(unsafe.Pointer(C.CreateStringFromString(ctx.ptr(), str.ptr())))
}

// Free a module string object obtained with one of the Redis modules API calls
// that return new string objects.
//
// It is possible to call this function even when automatic memory management
// is enabled. In that case the string will be released ASAP and removed
// from the pool of string to release at the end.
// void RM_FreeString(RedisModuleCtx *ctx, RedisModuleString *str);
func (ctx Ctx)FreeString(str String) () {
    C.FreeString(ctx.ptr(), str.ptr())
}

//
// int RM_WrongArity(RedisModuleCtx *ctx);
func (ctx Ctx)WrongArity() (int) {
    return int(C.WrongArity(ctx.ptr()))
}

// Send an integer reply to the client, with the specified long long value.
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithLongLong(RedisModuleCtx *ctx, long long ll);
func (ctx Ctx)ReplyWithLongLong(ll int64) (int) {
    return int(C.ReplyWithLongLong(ctx.ptr(), C.longlong(ll)))
}

// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithError(RedisModuleCtx *ctx, const char *err);
func (ctx Ctx)ReplyWithError(err string) (int) {
    c := C.CString(err)
    defer C.free(c);
    return int(C.ReplyWithError(ctx.ptr(), c))
}

// Reply with a simple string (+... \r\n in RESP protocol). This replies
// are suitable only when sending a small non-binary string with small
// overhead, like "OK" or similar replies.
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithSimpleString(RedisModuleCtx *ctx, const char *msg);
func (ctx Ctx)ReplyWithSimpleString(msg string) (int) {
    c := C.CString(msg)
    defer C.free(c);
    return int(C.ReplyWithSimpleString(ctx.ptr(), c))
}

// Reply with an array type of 'len' elements. However 'len' other calls
// to `ReplyWith*` style functions must follow in order to emit the elements
// of the array.
//
// When producing arrays with a number of element that is not known beforehand
// the function can be called with the special count
// `REDISMODULE_POSTPONED_ARRAY_LEN`, and the actual number of elements can be
// later set with `RedisModule_ReplySetArrayLength()` (which will set the
// latest "open" count if there are multiple ones).
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithArray(RedisModuleCtx *ctx, long len);
func (ctx Ctx)ReplyWithArray(len int64) (int) {
    return int(C.ReplyWithArray(ctx.ptr(), C.long(len)))
}

// When `RedisModule_ReplyWithArray()` is used with the argument
// `REDISMODULE_POSTPONED_ARRAY_LEN`, because we don't know beforehand the number
// of items we are going to output as elements of the array, this function
// will take care to set the array length.
//
// Since it is possible to have multiple array replies pending with unknown
// length, this function guarantees to always set the latest array length
// that was created in a postponed way.
//
// For example in order to output an array like [1,[10,20,30]] we
// could write:
//
//  `RedisModule_ReplyWithArray(ctx`,`REDISMODULE_POSTPONED_ARRAY_LEN`);
//  `RedisModule_ReplyWithLongLong(ctx`,1);
//  `RedisModule_ReplyWithArray(ctx`,`REDISMODULE_POSTPONED_ARRAY_LEN`);
//  `RedisModule_ReplyWithLongLong(ctx`,10);
//  `RedisModule_ReplyWithLongLong(ctx`,20);
//  `RedisModule_ReplyWithLongLong(ctx`,30);
//  `RedisModule_ReplySetArrayLength(ctx`,3); // Set len of 10,20,30 array.
//  `RedisModule_ReplySetArrayLength(ctx`,2); // Set len of top array
//
// Note that in the above example there is no reason to postpone the array
// length, since we produce a fixed number of elements, but in the practice
// the code may use an interator or other ways of creating the output so
// that is not easy to calculate in advance the number of elements.
// void RM_ReplySetArrayLength(RedisModuleCtx *ctx, long len);
func (ctx Ctx)ReplySetArrayLength(len int64) () {
    C.ReplySetArrayLength(ctx.ptr(), C.long(len))
}

// Reply with a bulk string, taking in input a C buffer pointer and length.
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithStringBuffer(RedisModuleCtx *ctx, const char *buf, size_t len);
func (ctx Ctx)ReplyWithStringBuffer(buf []byte, len int) (int) {
    // TODO free
    return int(C.ReplyWithStringBuffer(ctx.ptr(), (*C.char)(C.CBytes(buf)), C.size_t(len)))
}

// Reply with a bulk string, taking in input a RedisModuleString object.
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithString(RedisModuleCtx *ctx, RedisModuleString *str);
func (ctx Ctx)ReplyWithString(str String) (int) {
    return int(C.ReplyWithString(ctx.ptr(), str.ptr()))
}

// Reply to the client with a NULL. In the RESP protocol a NULL is encoded
// as the string "$-1\r\n".
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithNull(RedisModuleCtx *ctx);
func (ctx Ctx)ReplyWithNull() (int) {
    return int(C.ReplyWithNull(ctx.ptr()))
}

// Reply exactly what a Redis command returned us with `RedisModule_Call()`.
// This function is useful when we use `RedisModule_Call()` in order to
// execute some command, as we want to reply to the client exactly the
// same reply we obtained by the command.
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithCallReply(RedisModuleCtx *ctx, RedisModuleCallReply *reply);
func (ctx Ctx)ReplyWithCallReply(reply CallReply) (int) {
    return int(C.ReplyWithCallReply(ctx.ptr(), unsafe.Pointer(reply.ptr())))
}

// Send a string reply obtained converting the double 'd' into a bulk string.
// This function is basically equivalent to converting a double into
// a string into a C buffer, and then calling the function
// `RedisModule_ReplyWithStringBuffer()` with the buffer and length.
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplyWithDouble(RedisModuleCtx *ctx, double d);
func (ctx Ctx)ReplyWithDouble(d float64) (int) {
    return int(C.ReplyWithDouble(ctx.ptr(), C.double(d)))
}

// Replicate the specified command and arguments to slaves and AOF, as effect
// of execution of the calling command implementation.
//
// The replicated commands are always wrapped into the MULTI/EXEC that
// contains all the commands replicated in a given module command
// execution. However the commands replicated with `RedisModule_Call()`
// are the first items, the ones replicated with `RedisModule_Replicate()`
// will all follow before the EXEC.
//
// Modules should try to use one interface or the other.
//
// This command follows exactly the same interface of `RedisModule_Call()`,
// so a set of format specifiers must be passed, followed by arguments
// matching the provided format specifiers.
//
// Please refer to `RedisModule_Call()` for more information.
//
// The command returns `REDISMODULE_ERR` if the format specifiers are invalid
// or the command name does not belong to a known command.
// int RM_Replicate(RedisModuleCtx *ctx, const char *cmdname, const char *fmt, ...);
func (ctx Ctx)Replicate(cmdname string, format string, args ...interface{}) (int) {
    c := C.CString(cmdname)
    defer C.free(c);
    msg := fmt.Sprintf(format, args...)
    s := C.CString(msg)
    defer C.free(s);
    return int(C.Replicate(ctx.ptr(), c, s))
}

// This function will replicate the command exactly as it was invoked
// by the client. Note that this function will not wrap the command into
// a MULTI/EXEC stanza, so it should not be mixed with other replication
// commands.
//
// Basically this form of replication is useful when you want to propagate
// the command to the slaves and AOF file exactly as it was called, since
// the command can just be re-executed to deterministically re-create the
// new state starting from the old one.
//
// The function always returns `REDISMODULE_OK`.
// int RM_ReplicateVerbatim(RedisModuleCtx *ctx);
func (ctx Ctx)ReplicateVerbatim() (int) {
    return int(C.ReplicateVerbatim(ctx.ptr()))
}

// Return the ID of the current client calling the currently active module
// command. The returned ID has a few guarantees:
//
// 1. The ID is different for each different client, so if the same client
//    executes a module command multiple times, it can be recognized as
//    having the same ID, otherwise the ID will be different.
// 2. The ID increases monotonically. Clients connecting to the server later
//    are guaranteed to get IDs greater than any past ID previously seen.
//
// Valid IDs are from 1 to 2^64-1. If 0 is returned it means there is no way
// to fetch the ID in the context the function was currently called.
// unsigned long long RM_GetClientId(RedisModuleCtx *ctx);
func (ctx Ctx)GetClientId() (uint64) {
    return uint64(C.GetClientId(ctx.ptr()))
}

// Return the currently selected DB.
// int RM_GetSelectedDb(RedisModuleCtx *ctx);
func (ctx Ctx)GetSelectedDb() (int) {
    return int(C.GetSelectedDb(ctx.ptr()))
}

// Change the currently selected DB. Returns an error if the id
// is out of range.
//
// Note that the client will retain the currently selected DB even after
// the Redis command implemented by the module calling this function
// returns.
//
// If the module command wishes to change something in a different DB and
// returns back to the original one, it should call `RedisModule_GetSelectedDb()`
// before in order to restore the old DB number before returning.
// int RM_SelectDb(RedisModuleCtx *ctx, int newid);
func (ctx Ctx)SelectDb(newid int) (int) {
    return int(C.SelectDb(ctx.ptr(), C.int(newid)))
}

// Return an handle representing a Redis key, so that it is possible
// to call other APIs with the key handle as argument to perform
// operations on the key.
//
// The return value is the handle repesenting the key, that must be
// closed with `RM_CloseKey()`.
//
// If the key does not exist and WRITE mode is requested, the handle
// is still returned, since it is possible to perform operations on
// a yet not existing key (that will be created, for example, after
// a list push operation). If the mode is just READ instead, and the
// key does not exist, NULL is returned. However it is still safe to
// call `RedisModule_CloseKey()` and `RedisModule_KeyType()` on a NULL
// value.
// void *RM_OpenKey(RedisModuleCtx *ctx, robj *keyname, int mode);
func (ctx Ctx)OpenKey(keyname String, mode int) (unsafe.Pointer) {
    return unsafe.Pointer(C.OpenKey(ctx.ptr(), keyname.ptr(), C.int(mode)))
}


// Exported API to call any Redis command from modules.
// On success a RedisModuleCallReply object is returned, otherwise
// NULL is returned and errno is set to the following values:
//
// EINVAL: command non existing, wrong arity, wrong format specifier.
// EPERM:  operation in Cluster instance with key in non local slot.
// RedisModuleCallReply *RM_Call(RedisModuleCtx *ctx, const char *cmdname, const char *fmt, ...);
func (ctx Ctx)Call(cmdname string, format string, args ... interface{}) (CallReply) {
    c := C.CString(cmdname)
    defer C.free(c);
    msg := fmt.Sprintf(format, args...)
    s := C.CString(msg)
    defer C.free(s);
    return CreateCallReply(unsafe.Pointer(C.Call(ctx.ptr(), c, s)))
}


// Produces a log message to the standard Redis log, the format accepts
// printf-alike specifiers, while level is a string describing the log
// level to use when emitting the log, and must be one of the following:
//
// * "debug"
// * "verbose"
// * "notice"
// * "warning"
//
// If the specified log level is invalid, verbose is used by default.
// There is a fixed limit to the length of the log line this function is able
// to emit, this limti is not specified but is guaranteed to be more than
// a few lines of text.
// void RM_Log(RedisModuleCtx *ctx, const char *levelstr, const char *fmt, ...);
func (ctx Ctx)Log(l LogLevel, msg string) () {
    c := C.CString(msg)
    defer C.free(c);
    C.CtxLog(ctx.ptr(), C.int(l), c)
}

func (ctx Ctx)Init(name string, version int, apiVersion int) int {
    c := C.CString(name)
    defer C.free(c);
    return (int)(C.RedisModule_Init(ctx.ptr(), c, (C.int)(version), (C.int)(apiVersion)))
}
func (c Ctx)Load(mod *Module) int {
    if mod == nil {
        LogErr("Load Mod must not nil")
        return ERR
    }
    if c.Init(mod.Name, mod.Version, APIVER_1) == ERR {
        LogErr("Init mod %s failed", mod.Name)
        return ERR
    }
    for _, cmd := range mod.Commands {
        //if c.CreateCommand(cmd.Name, cmd.Action, "write fast deny-oom", 1, 1, 1) == ERR {
        //    LogErr("Create mod %s command %s failed", mod.Name, cmd.Name)
        //    return ERR
        //}
        LogDebug("Create mod %s command %s", mod.Name, cmd.Name)
    }
    return OK
}
//
//func (c Ctx)CreateCommand(name string, cmdFunc CmdFunc, strFlags string, firstKey int, lastKey int, keyStep int) int {
//    id := len(callbacks)
//    callbacks = append(callbacks, cmdFunc)
//    return (int)(C.CreateCommandCallID(c.ptr(), C.CString(name), C.int(id), C.CString(strFlags), C.int(firstKey), C.int(lastKey), C.int(keyStep)))
//}


// =============================================================================
// ========================== Reply functions
// =============================================================================


// Wrapper for the recursive free reply function. This is needed in order
// to have the first level function to return on nested replies, but only
// if called by the module API.
// void RM_FreeCallReply(RedisModuleCallReply *reply);
func (reply CallReply)FreeCallReply() () {
    C.FreeCallReply(reply.ptr())
}

// Return the reply type.
// int RM_CallReplyType(RedisModuleCallReply *reply);
func (reply CallReply)CallReplyType() (int) {
    return int(C.CallReplyType(reply.ptr()))
}

// Return the reply type length, where applicable.
// size_t RM_CallReplyLength(RedisModuleCallReply *reply);
func (reply CallReply)CallReplyLength() (int) {
    return int(C.CallReplyLength(reply.ptr()))
}

// Return the 'idx'-th nested call reply element of an array reply, or NULL
// if the reply type is wrong or the index is out of range.
// RedisModuleCallReply *RM_CallReplyArrayElement(RedisModuleCallReply *reply, size_t idx);
func (reply CallReply)CallReplyArrayElement(idx int) (CallReply) {
    return CreateCallReply(unsafe.Pointer(C.CallReplyArrayElement(reply.ptr(), C.size_t(idx))))
}

// Return the long long of an integer reply.
// long long RM_CallReplyInteger(RedisModuleCallReply *reply);
func (reply CallReply)CallReplyInteger() (int64) {
    return int64(C.CallReplyInteger(reply.ptr()))
}

// Return the pointer and length of a string or error reply.
// const char *RM_CallReplyStringPtr(RedisModuleCallReply *reply, size_t *len);
func (reply CallReply)CallReplyStringPtr(len *int) (unsafe.Pointer) {
    return unsafe.Pointer(C.CallReplyStringPtr(reply.ptr(), (*C.size_t)(unsafe.Pointer(len))))
}

// Return a new string object from a call reply of type string, error or
// integer. Otherwise (wrong reply type) return NULL.
// RedisModuleString *RM_CreateStringFromCallReply(RedisModuleCallReply *reply);
func (reply CallReply)CreateStringFromCallReply() (String) {
    return CreateString(unsafe.Pointer(C.CreateStringFromCallReply(reply.ptr())))
}

// Return a pointer, and a length, to the protocol returned by the command
// that returned the reply object.
// const char *RM_CallReplyProto(RedisModuleCallReply *reply, size_t *len);
func (reply CallReply)CallReplyProto(len *uint64) (unsafe.Pointer) {
    return unsafe.Pointer(C.CallReplyProto(reply.ptr(), (*C.size_t)(len)))
}

// =============================================================================
// ========================== String functions
// =============================================================================


// Given a string module object, this function returns the string pointer
// and length of the string. The returned pointer and length should only
// be used for read only accesses and never modified.
// const char *RM_StringPtrLen(RedisModuleString *str, size_t *len);
func (str String)StringPtrLen(len *uint64) (string) {
    return C.GoString(C.StringPtrLen(str.ptr(), (*C.size_t)(len)))
}

// Convert the string into a long long integer, storing it at `*ll`.
// Returns `REDISMODULE_OK` on success. If the string can't be parsed
// as a valid, strict long long (no spaces before/after), `REDISMODULE_ERR`
// is returned.
// int RM_StringToLongLong(RedisModuleString *str, long long *ll);
func (str String)StringToLongLong(ll *int64) (int) {
    return int(C.StringToLongLong(str.ptr(), (*C.longlong)(ll)))
}

// Convert the string into a double, storing it at `*d`.
// Returns `REDISMODULE_OK` on success or `REDISMODULE_ERR` if the string is
// not a valid string representation of a double value.
// int RM_StringToDouble(RedisModuleString *str, double *d);
func (str String)StringToDouble(d *float64) (int) {
    return int(C.StringToDouble(str.ptr(), (*C.double)(d)))
}

// =============================================================================
// ========================== Key functions
// =============================================================================

// =============================================================================
// ========================== IO functions
// =============================================================================