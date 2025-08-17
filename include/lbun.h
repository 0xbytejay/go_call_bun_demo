#ifndef LBUN_H
#define LBUN_H
#include <stdbool.h>
#include <stdint.h>
#ifdef __cplusplus
extern "C" {
#endif

typedef int64_t JSValue;
typedef void VirtualMachine;
typedef void JSGlobalObject;

typedef struct {
  JSValue value;
  bool is_err;
} JSValueResult;

typedef struct {
  uintptr_t data;
  uintptr_t len;
  uintptr_t cap;
} GoSliceHeader;

void startBunCli(GoSliceHeader *argv);

VirtualMachine *VirtualMachine_getMainThreadVM(void);
JSGlobalObject *VirtualMachine_getMainThreadVMGlobalObject(void);

JSValueResult JSGlobalObject_toJSValue(void *globalThis);
JSValueResult JSValue_get(void *globalThis, JSValue jsValue, char *property);
JSValueResult JSValue_get(void *globalThis, JSValue jsValue, char *property);
JSValueResult JSValue_fromInt64(void *globalThis, int64_t value);

char *JSValue_call(void *globalThis, JSValue thisValue, JSValue func,
                   GoSliceHeader *property);

char *JSValue_toString(void *globalThis, JSValue jsValue);



#ifdef __cplusplus
}
#endif

#endif /* LBUN_H */
