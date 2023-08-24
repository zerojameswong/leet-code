# brute force-ish solution beats 5%
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m = len(mat)
        n = len(mat[0])

        dist = [[-1 for j in range(n)] for i in range(m)]
        count = 0
        total_count = m * n

        for i in range(m):
            for j in range(n):
                if mat[i][j] == 0:
                    dist[i][j] = 0
                    count += 1

        current_dist = 0
        while count < total_count:
            for i in range(m):
                for j in range(n):
                    if dist[i][j] == current_dist:
                        if i - 1 >= 0 and dist[i - 1][j] == -1:
                            dist[i - 1][j] = current_dist + 1
                            count += 1
                        if j - 1 >= 0 and dist[i][j - 1] == -1:
                            dist[i][j - 1] = current_dist + 1
                            count += 1
                        if i + 1 < m and dist[i + 1][j] == -1:
                            dist[i + 1][j] = current_dist + 1
                            count += 1
                        if j + 1 < n and dist[i][j + 1] == -1:
                            dist[i][j + 1] = current_dist + 1
                            count += 1
            current_dist += 1

        return dist

# using queue.Queue for BFS, beats 12%
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m = len(mat)
        n = len(mat[0])
        total_count = m * n

        dist_mat = [[total_count for j in range(n)] for i in range(m)]

        import queue
        q = queue.Queue()

        for i in range(m):
            for j in range(n):
                if mat[i][j] == 0:
                    dist_mat[i][j] = 0
                    q.put((i, j))

        while not q.empty():
            i, j = q.get()
            dist = dist_mat[i][j]

            if i - 1 >= 0 and dist_mat[i - 1][j] > dist + 1:
                dist_mat[i - 1][j] = dist + 1
                q.put((i - 1, j))
            if j - 1 >= 0 and dist_mat[i][j - 1] > dist + 1:
                dist_mat[i][j - 1] = dist + 1
                q.put((i, j - 1))
            if i + 1 < m and dist_mat[i + 1][j] > dist + 1:
                dist_mat[i + 1][j] = dist + 1
                q.put((i + 1, j))
            if j + 1 < n and dist_mat[i][j + 1] > dist + 1:
                dist_mat[i][j + 1] = dist + 1
                q.put((i, j + 1))
                
        return dist_mat

# beats 93% by using collections.deque instead
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m = len(mat)
        n = len(mat[0])
        total_count = m * n

        dist_mat = [[total_count for j in range(n)] for i in range(m)]

        q = collections.deque()

        for i in range(m):
            for j in range(n):
                if mat[i][j] == 0:
                    dist_mat[i][j] = 0
                    q.append((i, j))

        while q:
            i, j = q.popleft()
            dist = dist_mat[i][j]

            if i - 1 >= 0 and dist_mat[i - 1][j] > dist + 1:
                dist_mat[i - 1][j] = dist + 1
                q.append((i - 1, j))
            if j - 1 >= 0 and dist_mat[i][j - 1] > dist + 1:
                dist_mat[i][j - 1] = dist + 1
                q.append((i, j - 1))
            if i + 1 < m and dist_mat[i + 1][j] > dist + 1:
                dist_mat[i + 1][j] = dist + 1
                q.append((i + 1, j))
            if j + 1 < n and dist_mat[i][j + 1] > dist + 1:
                dist_mat[i][j + 1] = dist + 1
                q.append((i, j + 1))
                
        return dist_mat
