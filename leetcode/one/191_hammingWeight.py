#! /usr/bin/env python

def hammingWeight(n):
    """
    :type n: int
    :rtype: int
    """
    cnt = 0
    length = 32
    while length > 0:
        if n & 1 == 1:
            cnt = cnt + 1
        n = n >> 1
        length = length - 1
    return cnt


def reverseBits(n):
    m, l = 0, 0
    while l < 32:
        print(l, m, n&1)
        m = m ^ (n&1)
        if l == 31:
            break
        m = m << 1
        n = n >> 1
        l = l + 1
    return m

if __name__ == '__main__':
    # print(hammingWeight(1))
    print(reverseBits(2147483648))