from math import sqrt

import numpy as np
import multiprocessing as mp
import time
import write_utils


def thread_mult(matrix1, matrix2, N, p_step_i, p_step_j, p_sqrt):
    result_matrix = np.zeros((N, N))
    # perform shift by process number
    # print(matrix1)
    for k in range(p_sqrt*p_sqrt):
        # matrix multiplication
        for i in range(p_step_i, p_step_i + p_sqrt):
            for j in range(p_step_j, p_step_j + p_sqrt):
                result_matrix[i][j] += matrix1[i][j] * matrix2[i][j]
        # perform shift by 1
        for i in range(N):
            # shift left
            matrix1[i] = np.roll(matrix1[i], -1)
        for i in range(N):
            # shift up
            matrix2[:, i] = np.roll(matrix2[:, i], -1)
    write_utils.write_conc_to_csv(int(N/p_sqrt), matrix1, matrix2, result_matrix)
    return result_matrix


def begin(matrix1, matrix2, P, N):
    write_utils.init_conc()
    start = time.time()
    result = np.zeros((N,N))
    # perform initial shift
    for i in range(N):
        # shift left
        matrix1[i] = np.roll(matrix1[i], -i)
        matrix2[:, i] = np.roll(matrix2[:, i], -i)

    p_sqrt = int(sqrt(P))
    n_p_sqrt = int(N/p_sqrt)
    # Init multiprocessing.Pool()
    pool = mp.Pool(mp.cpu_count())
    results = []
    for i in range(p_sqrt):
        for j in range(p_sqrt):
            results.append(pool.apply_async(thread_mult, args=(matrix1, matrix2, N, i*n_p_sqrt, j*n_p_sqrt, n_p_sqrt)))
    for i in range(P):
        result = np.add(result, results[i].get())
    pool.terminate()
    pool.close()
    pool.join()
    end = time.time()
    return end - start
