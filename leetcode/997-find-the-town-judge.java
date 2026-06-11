class Solution {
    public int findJudge(int n, int[][] trust) {
        int[] trustScores = new int[n+1];
        for (int[] relationShip : trust) {
            trustScores[relationShip[1]] += 1;
            trustScores[relationShip[0]] -= 1;
        }

        for (int i=1;i<n+1;i++) {
            if (trustScores[i] == n - 1) {
                return i;
            }
        }
        return -1;
    }
    public static void main(String[] args) {
	Solution sol = new Solution();
	System.out.println(sol.findJudge(2, new int[][]{{1,2}}));
	System.out.println(sol.findJudge(3, new int[][]{{1,3},{2,3}}));
    }
}

