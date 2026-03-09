#!/usr/bin/env python3
"""Huge Number Calculator - Python CLI"""

import sys


def add_big_numbers(num1: str, num2: str) -> str:
    """Add two large integers represented as strings."""
    # Remove leading zeros and handle signs
    num1 = num1.lstrip('0') or '0'
    num2 = num2.lstrip('0') or '0'

    # Reverse strings for easier digit-by-digit addition
    n1, n2 = num1[::-1], num2[::-1]

    result = []
    carry = 0
    max_len = max(len(n1), len(n2))

    for i in range(max_len):
        d1 = int(n1[i]) if i < len(n1) else 0
        d2 = int(n2[i]) if i < len(n2) else 0

        total = d1 + d2 + carry
        result.append(str(total % 10))
        carry = total // 10

    if carry:
        result.append(str(carry))

    return ''.join(result[::-1])


def main():
    if len(sys.argv) != 4 or sys.argv[2] != '+':
        print("Usage: python big_add.py <num1> + <num2>", file=sys.stderr)
        sys.exit(1)

    num1 = sys.argv[1].strip()
    num2 = sys.argv[3].strip()

    result = add_big_numbers(num1, num2)
    print(result)


if __name__ == "__main__":
    main()
