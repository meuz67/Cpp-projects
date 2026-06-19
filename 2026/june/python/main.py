def deleteRepeats(arr: []):
    seen = set()
    result = []
    for item in arr:
        if item not in seen:
            result.append(item)
            seen.add(item)
    return result
def main():
    print(deleteRepeats([1, 2, 3, 4, 4, 2]))

if __name__ == '__main__':
    main()