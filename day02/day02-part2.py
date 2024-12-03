def check_safe(report):
    
    def is_good_diff(d):
        return d > 0 and d <=3
    
    levels = [int(l) for l in report.split(" ")]
    diffs = [a - b for a, b in zip(levels[:-1], levels[1:])]
    
    # If levels are decresing, multiply by -1 to make them increase
    if sum(1 if d > 0 else -1 for d in diffs) < 0:
        diffs = [-d for d in diffs]
    
    good_diffs = [is_good_diff(d) for d in diffs]
    # simple case where all diffs are good
    if all(good_diffs):
        return True
    
    # if there are one or two bad diffs, check whether this can be fixed by removing a value
    if sum(good_diffs) >= len(diffs) - 2:
        false_indices = [i for i, val in enumerate(good_diffs) if not val]
        
         # If there are 2 false indices next to each other, check if their sum would be a good diff
        if len(false_indices) == 2 and abs(false_indices[1] - false_indices[0]) == 1:
            diff = diffs[false_indices[0]] + diffs[false_indices[1]]
            if is_good_diff(diff):
                return True
        
         # If there is 1 false index, check if 
        elif len(false_indices) == 1:
            false_index = false_indices[0]
            if false_index == 0 or false_index == len(diffs) - 1:
                return True
            diff1 = diffs[false_index] + diffs[false_index + 1]
            diff2 = diffs[false_index] + diffs[false_index - 1]
            if is_good_diff(diff1) or is_good_diff(diff2):
                return True

        return False
    
    return False


with open("input-ines.txt", "r") as f:
    reports = f.read().split("\n")

result = 0
for report in reports:
    if check_safe(report):
        result += 1

print(result)