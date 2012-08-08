
class hosting_packages::vcs {
    
    
    $vcs = ["git", "subversion","mercurial"]
    
    package { $vcs:
        ensure  => installed,
        require => Class["yumrepos::epel"],
        notify => Class["cloudlinux::cagefs_update"],
    }
    
}
