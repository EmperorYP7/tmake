#include <bits/stdc++.h>

using namespace std;

#define FASTIO ios_base::sync_with_stdio(false);\
               cin.tie(NULL);\
               cout.tie(NULL)
#define loop(i, a, b) for(fint i = a; i < (b); ++i)
#define rloop(i, a, b) for (fint i = b; i >= (a); --i)
#define TXTIO   freopen("input.txt", "r", stdin);\
                freopen("output.txt", "w", stdout)
#define print(x) cout << (x) << "\n"
#define read(x, n) loop(i, 0, n) { cin >> x[i]; }
#define printarr(x) for(auto it : x) { cout << it << " "; } cout << "\n"
#define pb push_back
#define abs(x) (((x) > 0) ? (x) : -(x))

typedef int_fast32_t i32;
typedef int_fast64_t fint;
typedef long long ll;
typedef long double ld;
typedef uint_fast64_t ufint;
typedef vector<fint> vi;
typedef vector<vector<fint>> vvi;
typedef unordered_map<fint, fint> mapii;
typedef pair<fint, fint> pii;
typedef priority_queue<fint> pqueue;
typedef set<fint> seti;
typedef multiset<fint> mseti;
typedef bitset<64> bits;

fint t, n;

void solve()
{
    cin >> n;
    
}

int32_t main()
{
#ifndef ONLINE_JUDGE
    TXTIO;
#endif
    FASTIO;
    cin>>t;

    while(t--)
    {
        solve();
    }

#ifndef ONLINE_JUDGE
    cout << "Time elapsed: " << 1000.0 * clock() / CLOCKS_PER_SEC << " ms\n";
#endif
}
