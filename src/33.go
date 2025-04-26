def reverse_string(s):
    """
    This function takes a string and returns its reverse.
    
    :param s: Input string
    :type s: str
    :return: Reverse of the input string
    :rtype: str
    """
    return s[::-1]

# Example usage
input_str = "hello"
output_str = reverse_string(input_str)
print(output_str)  # Output: "olleh"
