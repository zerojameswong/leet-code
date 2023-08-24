# not the most elegant solution but covers all the cases, beats 56%
# builds greedy line with spaces of 1, measured by current_line_length
# when the addition of a new word would exceed max_width, we process the accumulated words
# and put the new word into consideration for the next line
# the 3 cases for generating a new line are:
# 1. when there is a single word, just pad to the end of the word
# 2. when words can be evenly spaced, just join all words by that number of spaces
# 3. otherwise determine the number of spaces between each word using a recursive function
# finally we deal with the case that this is the last iteration and a line was already added
# and the final word is just chilling in the current_words list by padding to the right

class Solution:
    def get_space_counts(self, total_space_count, num_gaps, target, space_counts):
        if len(space_counts) == target:
            return space_counts

        current_space_count = total_space_count // num_gaps

        return self.get_space_counts(total_space_count - current_space_count, num_gaps - 1, target,[current_space_count] + space_counts)

    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        lines = []

        current_words = []
        current_words_length = 0
        current_line_length = 0

        for idx, word in enumerate(words):
            new_length = 0
            if current_line_length == 0:
                new_length = len(word)
            else:
                new_length = current_line_length + 1 + len(word)

            if new_length > maxWidth:
                remaining_space = maxWidth - current_words_length

                if len(current_words) == 1:
                    line = current_words[0] + ' ' * remaining_space
                    lines.append(line)

                elif remaining_space % (len(current_words) - 1) == 0:
                    num_spaces = remaining_space // (len(current_words) - 1)
                    line = (' ' * num_spaces).join(current_words)
                    lines.append(line)

                else:
                    space_counts = self.get_space_counts(remaining_space, len(current_words) - 1, len(current_words) - 1, [])
                    line = ''
                    for i in range(0, len(current_words) - 1):
                        line += current_words[i]
                        line += ' ' * space_counts[i]

                    line += current_words[-1]
                    lines.append(line)                        

                current_words = [word]
                current_line_length = len(word)
                current_words_length = len(word)
            else:
                current_words.append(word)
                current_line_length = new_length
                current_words_length += len(word)

            if idx == len(words) - 1:
                line = ' '.join(current_words)
                line += (maxWidth - len(line)) * ' '
                lines.append(line)

        return lines
    