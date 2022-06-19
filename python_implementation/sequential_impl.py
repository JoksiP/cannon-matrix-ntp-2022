import numpy as np
import time
import write_utils


def begin(matrix1, matrix2, result_matrix, P, M, L):
    write_utils.init_seq()
    start = time.time()
    # perform initial shift
    for i in range(M):
        # shift left
        matrix1[i] = np.roll(matrix1[i], -i)
    for i in range(L):
        # shift up
        matrix2[:, i] = np.roll(matrix2[:, i], -i)
    # P sequential iterations
    for k in range(P):
        # matrix multiplication
        for i in range(M):
            for j in range(L):
                result_matrix[i][j] += matrix1[i][j] * matrix2[i][j]
        # perform shift by 1
        for i in range(M):
            # shift left
            matrix1[i] = np.roll(matrix1[i], -1)
        for i in range(L):
            # shift up
            matrix2[:, i] = np.roll(matrix2[:, i], -1)
        write_utils.write_seq_to_csv(k, matrix1, matrix2, result_matrix)

    end = time.time()
    return end - start

