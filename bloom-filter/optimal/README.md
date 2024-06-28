In calculating the optimal parameters for a Bloom filter, the use of `math.Ceil` and `math.Round` functions ensures that the parameters are appropriately sized to meet the desired false positive probability and expected number of elements. Here's why each function is used:

### 1. Using `math.Ceil` for Bit Array Size (`m`)

The formula for calculating the bit array size `m` is:

\[ m = - \frac{n \cdot \ln(p)}{(\ln(2))^2} \]

This formula can produce a non-integer value, but `m` needs to be an integer because it represents the number of bits in the array. Since we can't have a fraction of a bit, we use `math.Ceil` to round up to the nearest whole number. Rounding up ensures that the bit array is large enough to achieve the desired false positive probability. If we were to round down, the bit array might be too small, leading to a higher false positive rate than specified.

### 2. Using `math.Round` for Number of Hash Functions (`k`)

The formula for calculating the number of hash functions `k` is:

\[ k = \frac{m}{n} \cdot \ln(2) \]

This formula also can produce a non-integer value, but `k` needs to be an integer because it represents the number of hash functions. We use `math.Round` to round to the nearest integer to get the closest approximation. Rounding ensures that we don't systematically bias the number of hash functions up or down. If we were to always round up or down, we might either overuse hash functions (leading to higher computational cost) or underuse them (leading to a higher false positive rate).


### Creating Optimal Bloom filters

#### Line: `bitSet := make([]byte, (size+7)/8)`
- **Purpose**: To allocate a byte slice (`bitSet`) to store the bits for the Bloom filter.
- **Reasoning**: Each byte can hold 8 bits. To store `size` bits, we need `(size + 7) / 8` bytes. The `+7` ensures that any extra bits are rounded up to the next byte.

#### Line: `bf.bitSet[idx/8] |= 1 << (idx % 8)`
- **Purpose**: To set a specific bit in the `bitSet` array.
- **Reasoning**:
  - `idx / 8` gives the byte index in `bitSet`.
  - `1 << (idx % 8)` creates a bitmask with a 1 at the bit position `idx % 8` within that byte.
  - `|=` sets the bit at that position without affecting other bits in the byte.

#### Line: `if bf.bitSet[idx/8]&(1<<(idx%8)) == 0`
- **Purpose**: To check if a specific bit in the `bitSet` array is set.
- **Reasoning**:
  - `idx / 8` gives the byte index in `bitSet`.
  - `1 << (idx % 8)` creates a bitmask with a 1 at the bit position `idx % 8` within that byte.
  - `&` checks if that bit is set.
  - If the result is `0`, the bit is not set, meaning the element is definitely not in the set.


### Summary

- **`math.Ceil`**: Used to round up the bit array size `m` to ensure we have enough bits to meet the desired false positive probability.
- **`math.Round`**: Used to round the number of hash functions `k` to the nearest integer to get the closest approximation without systematic bias.