import csv


def init_seq():
    with open('seq_outputs.csv', 'w', encoding='UTF8') as f:
        writer = csv.writer(f)

        # write the header
        writer.writerow(['Iteration No.', 'Matrix 1', 'Matrix 2', 'Result'])


def write_seq_to_csv(iteration, m1, m2, result):
    with open('seq_outputs.csv', 'a+', encoding='UTF8') as f:
        writer = csv.writer(f)

        # write the data
        writer.writerow([iteration, m1, m2, result])

def init_conc():
    with open('conc_outputs.csv', 'w', encoding='UTF8') as f:
        writer = csv.writer(f)

        # write the header
        writer.writerow(['Process No.', 'Matrix 1', 'Matrix 2', 'Result'])


def write_conc_to_csv(iteration, m1, m2, result):
    with open('conc_outputs.csv', 'a+', encoding='UTF8') as f:
        writer = csv.writer(f)

        # write the data
        writer.writerow([iteration, m1, m2, result])
