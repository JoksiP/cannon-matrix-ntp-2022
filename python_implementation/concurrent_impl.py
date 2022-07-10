from math import sqrt

import numpy as np
import multiprocessing as mp
import time
import write_utils

def shift_parallel(matrix1, matrix2, N, step):
    for i in range(N):
        s = 1
        if step != 1:
            s = i
        # shift left
        matrix1[i] = np.roll(matrix1[i], -s)
        # shift up
        matrix2[:, i] = np.roll(matrix2[:, i], -s)
    return [matrix1, matrix2]


def mult_parallel(matrix1, matrix2, row, col):
    N = len(matrix1)
    res = np.zeros((N,N))
    for i in range(N):
        for j in range(N):
            res[i][j] += matrix1[i][j]*matrix2[i][j]
    return [res, row, col]


def begin(matrix1, matrix2, N, P):
    write_utils.init_conc()

    p_sqrt = int(sqrt(P))
    n_p_sqrt = int(N / p_sqrt)
    elapsed = time.time()
    result = np.zeros((N, N))

    thread_results = []
    # Init multiprocessing.Pool()
    pool = mp.Pool(mp.cpu_count())
    # perform initial shift
    proc_c = pool.apply_async(shift_parallel, args=(matrix1, matrix2, N, 2)).get()
    matrix1, matrix2 = proc_c[0], proc_c[1]

    for _ in range(N):
        for i in range(p_sqrt):
            for j in range(p_sqrt):
                row = n_p_sqrt * i
                col = n_p_sqrt * j
                m1 = matrix1[row:(row + n_p_sqrt), col:(col+n_p_sqrt)]
                m2 = matrix2[row:(row + n_p_sqrt), col:(col + n_p_sqrt)]

                thread_results.append(pool.apply_async(mult_parallel, args=(m1, m2, row, col)))
        matrix1, matrix2 = pool.apply_async(shift_parallel, args=(matrix1, matrix2, N, 1)).get()
    for i in range(len(thread_results)):
        thread_result = thread_results[i].get()
        row = thread_result[1]
        col = thread_result[2]
        for k in range(row, row + n_p_sqrt):
            for l in range(col, col + n_p_sqrt):
                result[k][l] += thread_result[0][k - row][l - col]
    pool.terminate()
    pool.close()
    pool.join()
    print(result)
    return time.time() - elapsed
