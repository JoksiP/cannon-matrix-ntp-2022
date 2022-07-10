import numpy as np
import sequential_impl
import concurrent_impl
import copy
import sys
from math import sqrt


def main():
    N = int(sys.argv[1])
    P = int(sys.argv[2])

    matrix1 = np.random.randint(1, 10, size=(N, N))
    matrix2 = np.random.randint(1, 10, size=(N, N))
    print(N)
    print(P)
    if int(sqrt(P)+0.5) ** 2 != P:
        print("P must be perfect square.")
        sys.exit()
    if int(N/sqrt(P))*int(sqrt(P)) != N:
        print("N must be dividable by the square root of P.")
        sys.exit()

    elapsed_seq = sequential_impl.begin(copy.deepcopy(matrix1), copy.deepcopy(matrix2), N, N)
    elapsed_conc = concurrent_impl.begin(copy.deepcopy(matrix1), copy.deepcopy(matrix2), N, P)
    print("Sequential implementation elapsed time: {} \nConcurrent implementation elapsed time: {}".format(elapsed_seq, elapsed_conc))


if __name__ == '__main__':
    main()
