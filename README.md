# AMockGen
AMockGen is a code generator for AMock.

# Overview
AMockGen generates a mock implementation (from `MockImplDesc`) of the interface.

With help of this mock implementation you can register/unregister methods calls.
Each method call is simply a function, which will be executed by calling its
corresponding interface method.
