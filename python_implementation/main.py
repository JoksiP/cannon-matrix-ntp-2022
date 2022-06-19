import numpy as np
import sequential_impl
import concurrent_impl
import copy


def main():
    # m x n
    matrix1 = np.random.randint(1, 10, size=(200, 200))
    # n x l
    matrix2 = np.random.randint(1, 10, size=(200, 200))
    # m x l
    M = len(matrix1)
    N = len(matrix2)
    L = len(matrix2[0])
    result_matrix = np.zeros((M, L))
    # number of processes
    P = N
    print(np.matmul(matrix1, matrix2))
    elapsed = sequential_impl.begin(copy.deepcopy(matrix1), copy.deepcopy(matrix2), result_matrix, P, M, L)
    print(elapsed)
    result_matrix = np.zeros((M, L))
    elapsed = concurrent_impl.begin(matrix1, matrix2, P, M, L)
    print(elapsed)


if __name__ == '__main__':
    main()