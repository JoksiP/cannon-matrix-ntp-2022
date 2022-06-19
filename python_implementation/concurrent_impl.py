import numpy as np
import multiprocessing as mp
import time
import write_utils
import copy

# Each thread / process
def conc_process(matrix1, matrix2, k_proc, M, L):
    result_matrix = np.zeros((M, L))
    # perform shift by process number
    #print(matrix1)
    for i in range(M):
        # shift left
        matrix1[i] = np.roll(matrix1[i], -k_proc)
    for i in range(L):
        # shift up
        matrix2[:, i] = np.roll(matrix2[:, i], -k_proc)
    # matrix multiplication
    for i in range(M):
        for j in range(L):
            result_matrix[i][j] += matrix1[i][j] * matrix2[i][j]
    write_utils.write_conc_to_csv(k_proc, matrix1, matrix2, result_matrix)
    return result_matrix

"""# Thread callback
def collect_result(result):
    global results_matrix
    results_matrix += result"""


# Main body
def begin(matrix1, matrix2, P, M, L):
    results_matrix = np.zeros((M, L))
    write_utils.init_conc()
    start = time.time()
    # perform initial shift
    for i in range(M):
        # shift left
        matrix1[i] = np.roll(matrix1[i], -i)
    for i in range(L):
        # shift up
        matrix2[:, i] = np.roll(matrix2[:, i], -i)

    # Init multiprocessing.Pool()
    pool = mp.Pool(mp.cpu_count())
    results = []
    for k in range(P):
        #m1 = copy.deepcopy(matrix1)
        #m2 = copy.deepcopy(matrix2)
        results.append(pool.apply_async(conc_process, args=(matrix1, matrix2, k, M, L)))
    for i in range(P):
        results_matrix = np.add(results_matrix, results[i].get())
    pool.terminate()
    pool.close()
    pool.join()
    print(results_matrix)
    return time.time() - start
