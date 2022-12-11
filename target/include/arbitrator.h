#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

#define ARBITRATOR_MACHINE_STATUS_RUNNING 0

#define ARBITRATOR_MACHINE_STATUS_FINISHED 1

#define ARBITRATOR_MACHINE_STATUS_ERRORED 2

#define ARBITRATOR_MACHINE_STATUS_TOO_FAR 3

#define GLOBAL_STATE_BYTES32_NUM 2

#define GLOBAL_STATE_U64_NUM 2

#define Machine_MAX_STEPS (1 << 43)

typedef struct Machine Machine;

typedef struct CByteArray {
  const uint8_t *ptr;
  uintptr_t len;
} CByteArray;

typedef struct Bytes32 {
  uint8_t bytes[32];
} Bytes32;

typedef struct GlobalState {
  struct Bytes32 bytes32_vals[GLOBAL_STATE_BYTES32_NUM];
  uint64_t u64_vals[GLOBAL_STATE_U64_NUM];
} GlobalState;

typedef struct ResolvedPreimage {
  uint8_t *ptr;
  intptr_t len;
} ResolvedPreimage;

typedef struct RustByteArray {
  uint8_t *ptr;
  uintptr_t len;
  uintptr_t capacity;
} RustByteArray;

struct Machine *arbitrator_load_machine(const char *binary_path,
                                        const char *const *library_paths,
                                        intptr_t library_paths_size);

struct Machine *arbitrator_load_wavm_binary(const char *binary_path);

void arbitrator_free_machine(struct Machine *mach);

struct Machine *arbitrator_clone_machine(struct Machine *mach);

/**
 * Go doesn't have this functionality builtin for whatever reason. Uses relaxed ordering.
 */
void atomic_u8_store(uint8_t *ptr, uint8_t contents);

/**
 * Runs the machine while the condition variable is zero. May return early if num_steps is hit.
 * Returns a c string error (freeable with libc's free) on error, or nullptr on success.
 */
char *arbitrator_step(struct Machine *mach, uint64_t num_steps, const uint8_t *condition);

int arbitrator_add_inbox_message(struct Machine *mach,
                                 uint64_t inbox_identifier,
                                 uint64_t index,
                                 struct CByteArray data);

/**
 * Like arbitrator_step, but stops early if it hits a host io operation.
 * Returns a c string error (freeable with libc's free) on error, or nullptr on success.
 */
char *arbitrator_step_until_host_io(struct Machine *mach, const uint8_t *condition);

int arbitrator_serialize_state(const struct Machine *mach, const char *path);

int arbitrator_deserialize_and_replace_state(struct Machine *mach, const char *path);

uint64_t arbitrator_get_num_steps(const struct Machine *mach);

/**
 * Returns one of ARBITRATOR_MACHINE_STATUS_*
 */
uint8_t arbitrator_get_status(const struct Machine *mach);

struct GlobalState arbitrator_global_state(struct Machine *mach);

void arbitrator_set_global_state(struct Machine *mach, struct GlobalState gs);

void arbitrator_set_preimage_resolver(struct Machine *mach,
                                      struct ResolvedPreimage (*resolver)(uint64_t, const uint8_t*));

void arbitrator_set_context(struct Machine *mach, uint64_t context);

struct Bytes32 arbitrator_hash(struct Machine *mach);

struct Bytes32 arbitrator_module_root(struct Machine *mach);

struct RustByteArray arbitrator_gen_proof(struct Machine *mach);

void arbitrator_free_proof(struct RustByteArray proof);
