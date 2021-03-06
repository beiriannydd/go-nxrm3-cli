// create-privilege.groovy is a  Nexus3 Integration API definition to create a repository privilege in Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// import nexus libraries from the nexus-security jar file
import org.sonatype.nexus.security.authz.AuthorizationManager
import org.sonatype.nexus.security.privilege.Privilege

// input map
Map input = new JsonSlurper().parseText(args)
// output map
Map output = [:]

authManager = container.lookup(AuthorizationManager)

Map properties = [:]
properties.put('contentSelector', input.properties.contentSelector)
properties.put('repository', input.properties.repository)
properties.put('actions', input.properties.actions)

privilege = new Privilege(
        id: input.id,
        name: input.name,
        description: input.description,
        type: input.type,
        version: '',
        properties: properties
)

authManager.addPrivilege(privilege)

// output success status
output.put('status', '200 OK')

// nexus logger
log.info('***********************************************')
log.info(String.format('Privilege %s is created', input.name))
log.info('**********************************************')

// return output in JSON format
JsonOutput.toJson(output)
