# Optimal Bloom Filter

A Bloom filter is a space-efficient probabilistic data structure designed to test whether an element is a member of a set. This implementation calculates the optimal parameters for the Bloom filter and uses bit manipulation techniques to efficiently manage memory and operations.

Bloom filters use hash functions to map elements to positions in a bit array, and MurmurHash is one of the popular choices for several reasons:

1. **Speed**:
    - MurmurHash is designed to be fast. It can process large amounts of data quickly, which is essential for applications like Bloom filters that require multiple hash computations for each insertion and lookup.

2. **Uniform Distribution**:
    - MurmurHash produces a uniform distribution of hash values. This property ensures that the bits in the Bloom filter are set and checked evenly, which helps in maintaining the desired false positive rate. Uniform distribution prevents clustering of bits, which would otherwise lead to increased false positives.

3. **Non-Cryptographic Hash Function**:
    - MurmurHash is a non-cryptographic hash function, meaning it is not designed to be secure against adversaries. However, for Bloom filters, the primary requirements are speed and distribution uniformity rather than cryptographic security.

4. **Deterministic Output**:
    - MurmurHash is deterministic, meaning it will always produce the same output for the same input. This property is crucial for Bloom filters, as the same element needs to hash to the same positions in the bit array every time it is processed.

5. **Low Collision Rate**:
    - MurmurHash has a low collision rate compared to simpler hash functions like modulo or division. Collisions (where two different inputs produce the same hash value) can increase the false positive rate in a Bloom filter. MurmurHash's design minimizes this risk, ensuring better accuracy for the filter.

6. **Flexibility**:
    - MurmurHash can produce 32-bit, 64-bit, or 128-bit hash values, offering flexibility depending on the requirements of the Bloom filter implementation. The 128-bit version, for example, provides a very large space to reduce collisions further.

7. **Simplicity**:
    - The implementation of MurmurHash is relatively simple and does not require complex operations or large libraries, making it easy to integrate into various applications, including Bloom filters.

### Example of Usage in Bloom Filter

In a Bloom filter, multiple hash functions are typically used. One way to generate multiple hash functions efficiently is to use a single base hash function (like MurmurHash) and modify it slightly (e.g., using different seeds) to simulate multiple independent hash functions. This is both computationally efficient and effective in maintaining a low false positive rate.

Here's a brief summary of why MurmurHash is favored in Bloom filters:

- **Speed**: Fast processing of elements.
- **Uniform Distribution**: Even spread of bits in the array.
- **Deterministic**: Consistent output for the same input.
- **Low Collision Rate**: Reduced false positives.
- **Flexibility**: Various bit-length outputs.
- **Simplicity**: Easy to implement and use.

These characteristics make MurmurHash an excellent choice for the hash functions used in Bloom filters.

## Memory Efficiency of byte Array vs. bool Array
Using a byte array instead of a bool array can save a significant amount of memory. Here’s how:

-* A bool in Go typically takes up 1 byte (8 bits) of memory.
-* A byte also takes up 1 byte (8 bits) of memory, but it can represent 8 different boolean values (one per bit).

So, if we use a bool array of size m, it will take up m bytes. In contrast, using a byte array of size m/8 will take up m/8 bytes, thus reducing the memory usage by a factor of 8.

## Overview

In calculating the optimal parameters for a Bloom filter, the use of `math.Ceil` and `math.Round` functions ensures that the parameters are appropriately sized to meet the desired false positive probability and expected number of elements. Here's why each function is used:

### Using `math.Ceil` for Bit Array Size (`m`)

The formula for calculating the bit array size `m` is:

\[ m = - \frac{n \cdot \ln(p)}{(\ln(2))^2} \]

This formula can produce a non-integer value, but `m` needs to be an integer because it represents the number of bits in the array. Since we can't have a fraction of a bit, we use `math.Ceil` to round up to the nearest whole number. Rounding up ensures that the bit array is large enough to achieve the desired false positive probability. If we were to round down, the bit array might be too small, leading to a higher false positive rate than specified.

### Using `math.Round` for Number of Hash Functions (`k`)

The formula for calculating the number of hash functions `k` is:

\[ k = \frac{m}{n} \cdot \ln(2) \]

This formula also can produce a non-integer value, but `k` needs to be an integer because it represents the number of hash functions. We use `math.Round` to round to the nearest integer to get the closest approximation. Rounding ensures that we don't systematically bias the number of hash functions up or down. If we were to always round up or down, we might either overuse hash functions (leading to higher computational cost) or underuse them (leading to a higher false positive rate).

## Creating Optimal Bloom Filters

### Allocating Memory for the Bit Array

```go
bitSet := make([]byte, (size+7)/8)
```

- **Purpose**: To allocate a byte slice (`bitSet`) to store the bits for the Bloom filter.
- **Reasoning**: Each byte can hold 8 bits. To store `size` bits, we need `(size + 7) / 8` bytes. The `+7` ensures that any extra bits are rounded up to the next byte.

### Setting a Bit in the Bit Array

```go
bf.bitSet[idx/8] |= 1 << (idx % 8)
```

- **Purpose**: To set a specific bit in the `bitSet` array.
- **Reasoning**:
  - `idx / 8` gives the byte index in `bitSet`.
  - `1 << (idx % 8)` creates a bitmask with a 1 at the bit position `idx % 8` within that byte.
  - `|=` sets the bit at that position without affecting other bits in the byte.

### Checking if a Bit is Set in the Bit Array

```go
if bf.bitSet[idx/8]&(1<<(idx%8)) == 0
```

- **Purpose**: To check if a specific bit in the `bitSet` array is set.
- **Reasoning**:
  - `idx / 8` gives the byte index in `bitSet`.
  - `1 << (idx % 8)` creates a bitmask with a 1 at the bit position `idx % 8` within that byte.
  - `&` checks if that bit is set.
  - If the result is `0`, the bit is not set, meaning the element is definitely not in the set.

## Trade-Offs in Tuning Parameters

When tuning the parameters `n` (the expected number of elements) and `p` (the desired false positive probability) in a Bloom filter, there are several trade-offs to consider. These trade-offs affect the memory usage, the number of hash functions, and the computational efficiency.

### Memory Usage (Size of the Bit Array `m`)

#### Trade-Off:

- **Lower `p` (lower false positive probability)**: Requires a larger bit array `m` to store the elements. This increases memory usage.
- **Higher `p` (higher false positive probability)**: Allows for a smaller bit array `m`, reducing memory usage but increasing the chance of false positives.

#### Formula:

\[ m = - \frac{n \cdot \ln(p)}{(\ln(2))^2} \]

### Number of Hash Functions `k`

#### Trade-Off:

- **Lower `p` (lower false positive probability)**: Typically requires more hash functions `k`. More hash functions increase the computational cost for adding and checking elements.
- **Higher `p` (higher false positive probability)**: Requires fewer hash functions, reducing computational overhead but increasing the chance of false positives.

#### Formula:

\[ k = \frac{m}{n} \cdot \ln(2) \]

### Computational Efficiency

#### Trade-Off:

- **More Hash Functions**: Increases the time complexity of both the `Add` and `Contains` operations. Each operation needs to compute multiple hash functions.
- **Fewer Hash Functions**: Decreases the time complexity but may increase the false positive rate.

### Practical Considerations

- **Memory Constraints**: If memory is a constraint, you may need to tolerate a higher false positive rate (`p`). This allows the Bloom filter to use less memory.
- **Performance Constraints**: If the performance is critical and you need faster `Add` and `Contains` operations, you might opt for fewer hash functions. This may require accepting a higher false positive rate.
- **Accuracy Requirements**: If accuracy (i.e., a very low false positive rate) is crucial for your application, you will need to allocate more memory for a larger bit array and use more hash functions.

## Examples

### Example 1: Low False Positive Rate

- **`n` = 1000** (expected elements)
- **`p` = 0.001** (desired false positive rate)

This configuration will result in a larger bit array (`m`) and more hash functions (`k`), leading to higher memory usage and computational cost but a very low false positive rate.

### Example 2: Higher False Positive Rate

- **`n` = 1000** (expected elements)
- **`p` = 0.01** (desired false positive rate)

This configuration will result in a smaller bit array (`m`) and fewer hash functions (`k`), leading to lower memory usage and computational cost but a higher false positive rate.

## Summary

- **Memory vs. False Positives**: Lowering `p` reduces false positives but increases memory usage.
- **Speed vs. Accuracy**: Fewer hash functions speed up operations but increase the false positive rate.
- **Memory and Performance Constraints**: Depending on your application's requirements, you need to balance between memory usage, computational efficiency, and the acceptable false positive rate.

By understanding these trade-offs, you can tune the parameters of your Bloom filter to best suit your application's needs.

